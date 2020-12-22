package conmap

import (
	"sort"
	"sync"
)

const (
	// Our concurrent hashmap has been intended to be internally divided into 32 shards,
	// so there will be at max 32 threads can concurrently access at the same time.
	__Shards = 32
)

type (
	// Shard defines the minimal segment that holds the golang map.
	Shard struct {
		rwl     sync.RWMutex
		backend map[string]interface{}
	}

	// ConMap is a kind of thread-safe map supporting concurrent reads and writes.
	ConMap struct {
		mu     sync.Mutex
		shards []*Shard
	}
)

// NewConMap returns a new ConMap instance.
func NewConMap() *ConMap {
	cm := ConMap{
		shards: make([]*Shard, __Shards),
	}
	for i := 0; i < __Shards; i++ {
		cm.shards[i] = &Shard{
			backend: make(map[string]interface{}),
		}
	}
	return &cm
}

func (cm *ConMap) lookup(k string) *Shard {
	return cm.shards[fnv_1a_32(k)&0x1f]
}

// Load retrieves a value relevant to the given key.
func (cm *ConMap) Load(k string) (v interface{}, ok bool) {
	s := cm.lookup(k)
	s.rwl.RLock()
	v, ok = s.backend[k]
	s.rwl.RUnlock()
	return v, ok
}

// BatchStore sets the given map.
func (cm *ConMap) BatchStore(m map[string]interface{}) {
	for k, v := range m {
		cm.Store(k, v)
	}
}

// Store sets the given value relevant to the given key.
func (cm *ConMap) Store(k string, v interface{}) {
	s := cm.lookup(k)
	s.rwl.Lock()
	s.backend[k] = v
	s.rwl.Unlock()
}

// StoreIfNotExists sets the given value relevant to the given key if no value associates with it.
func (cm *ConMap) StoreIfNotExists(k string, v interface{}) (ok bool) {
	s := cm.lookup(k)
	s.rwl.Lock()
	_, ok = s.backend[k]
	if !ok {
		s.backend[k] = v
	}
	s.rwl.Unlock()
	return !ok
}

// Remove deletes the given key.
func (cm *ConMap) Remove(k string) {
	s := cm.lookup(k)
	s.rwl.Lock()
	delete(s.backend, k)
	s.rwl.Unlock()
}

// Has checks if a given key existed.
func (cm *ConMap) Has(k string) (ok bool) {
	s := cm.lookup(k)
	s.rwl.RLock()
	_, ok = s.backend[k]
	s.rwl.RUnlock()
	return ok
}

// Count returns the total num of items inside the map.
func (cm *ConMap) Count() (cnt int) {
	cm.mu.Lock()
	cnt = 0
	for i := 0; i < __Shards; i++ {
		s := cm.shards[i]
		s.rwl.RLock()
		cnt += len(s.backend)
		s.rwl.RUnlock()
	}
	cm.mu.Unlock()
	return cnt
}

// Empty checks if map is empty.
func (cm *ConMap) Empty() bool {
	return cm.Count() == 0
}

// Item is used to wrap k/v together.
type Item struct {
	Key string
	Val interface{}
}

// Iter returns a buffered iterator which could be used in a for range loop.
func (cm *ConMap) Iter() (in chan *Item) {
	ous := cm.snapshot()
	in = make(chan *Item, 32)
	go fanIn(ous, in)
	return in
}

func (cm *ConMap) snapshot() (ous []chan *Item) {
	ous = make([]chan *Item, __Shards)
	for idx := range cm.shards {
		ous[idx] = make(chan *Item, 32)
		go func(ou chan *Item, s *Shard) {
			s.rwl.RLock()
			for k, v := range s.backend {
				ou <- &Item{Key: k, Val: v}
			}
			s.rwl.RUnlock()
			close(ou)
		}(ous[idx], cm.shards[idx])
	}
	return ous
}

func fanIn(ins []chan *Item, ou chan *Item) {
	wg := new(sync.WaitGroup)
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in chan *Item) {
			defer wg.Done()
			for p := range in {
				ou <- p
			}
		}(in)
	}
	wg.Wait()
	close(ou)
}

// Map returns all items as map[string]interface{}.
func (cm *ConMap) Map() (m map[string]interface{}) {
	m = make(map[string]interface{})
	for p := range cm.Iter() {
		m[p.Key] = p.Val
	}
	return m
}

// Keys returns all keys.
func (cm *ConMap) Keys(sorted bool) (keys []string) {
	keys = make([]string, cm.Count())
	idx := 0
	for p := range cm.Iter() {
		keys[idx] = p.Key
		idx++
	}
	if sorted {
		sort.Strings(keys)
	}
	return keys
}
