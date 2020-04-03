package comap

import (
	"sort"
	"sync"
)

const (
	// CoMap is internally divided into 32 segments, so at max 32 threads can concurrently access at a time.
	_Shards = 32
)

type (
	// Shard defines the minimal segment.
	Shard struct {
		mu      sync.RWMutex
		backend map[string]interface{}
	}

	// CoMap is a thread-safe map supporting concurrent reads and writes.
	CoMap []*Shard
)

// NewCoMap returns a new CoMap instance.
func NewCoMap() CoMap {
	cm := make(CoMap, _Shards)
	for i := 0; i < _Shards; i++ {
		cm[i] = &Shard{
			backend: make(map[string]interface{}),
		}
	}
	return cm
}

// Lookup returns the shard relevant to the given key.
func (cm CoMap) Lookup(k string) *Shard {
	return cm[fnv_1a_32(k)&0x1f]
}

// Load retrieves a value relevant to the given key.
func (cm CoMap) Load(k string) (v interface{}, ok bool) {
	s := cm.Lookup(k)
	s.mu.RLock()
	v, ok = s.backend[k]
	s.mu.RUnlock()
	return v, ok
}

// BatchStore sets the given map.
func (cm CoMap) BatchStore(m map[string]interface{}) {
	for k, v := range m {
		cm.Store(k, v)
	}
}

// Store sets the given value relevant to the given key.
func (cm CoMap) Store(k string, v interface{}) {
	s := cm.Lookup(k)
	s.mu.Lock()
	s.backend[k] = v
	s.mu.Unlock()
}

// StoreIfNotExists sets the given value relevant to the given key if no value associates with it.
func (cm CoMap) StoreIfNotExists(k string, v interface{}) (ok bool) {
	s := cm.Lookup(k)
	s.mu.Lock()
	_, ok = s.backend[k]
	if !ok {
		s.backend[k] = v
	}
	s.mu.Unlock()
	return !ok
}

// Remove deletes the given key.
func (cm CoMap) Remove(k string) {
	s := cm.Lookup(k)
	s.mu.Lock()
	delete(s.backend, k)
	s.mu.Unlock()
}

// Has checks if a given key existed.
func (cm CoMap) Has(k string) (ok bool) {
	s := cm.Lookup(k)
	s.mu.RLock()
	_, ok = s.backend[k]
	s.mu.RUnlock()
	return ok
}

// Count returns the total num of items inside the comap.
func (cm CoMap) Count() (cnt int) {
	cnt = 0
	for i := 0; i < _Shards; i++ {
		s := cm[i]
		s.mu.RLock()
		cnt += len(s.backend)
		s.mu.RUnlock()
	}
	return cnt
}

// Empty checks if comap is empty.
func (cm CoMap) Empty() bool {
	return cm.Count() == 0
}

// Item is used to wrap k/v together.
type Item struct {
	key string
	val interface{}
}

// Iter returns a buffered iterator which could be used in a for range loop.
func (cm CoMap) Iter() (ch chan *Item) {
	chs := cm.snapshot()
	ch = make(chan *Item, 32)
	go fanIn(chs, ch)
	return ch
}

func (cm CoMap) snapshot() (chs []chan *Item) {
	chs = make([]chan *Item, _Shards)
	for idx := range cm {
		chs[idx] = make(chan *Item, 32)
		go func(ch chan *Item, s *Shard) {
			s.mu.RLock()
			for k, v := range s.backend {
				ch <- &Item{key: k, val: v}
			}
			s.mu.RUnlock()
			close(ch)
		}(chs[idx], (cm)[idx])
	}
	return chs
}

func fanIn(ins []chan *Item, out chan *Item) {
	wg := new(sync.WaitGroup)
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in chan *Item) {
			defer wg.Done()
			for p := range in {
				out <- p
			}
		}(in)
	}
	wg.Wait()
	close(out)
}

// Map returns all items as map[string]interface{}.
func (cm CoMap) Map() (m map[string]interface{}) {
	m = make(map[string]interface{})
	for p := range cm.Iter() {
		m[p.key] = p.val
	}
	return m
}

// Keys returns all keys.
func (cm CoMap) Keys(sorted bool) (keys []string) {
	keys = make([]string, cm.Count())
	idx := 0
	for p := range cm.Iter() {
		keys[idx] = p.key
		idx++
	}
	if sorted {
		sort.Strings(keys)
	}
	return keys
}
