package conmap

import (
	"fmt"
	"sync"
	"testing"
)

type TestUsageMap struct {
	mu sync.Mutex
	m  map[string]string
}

func BenchmarkConMapThroughputBatch_1(b *testing.B) {
	cm := NewConMap()
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		for j := 0; j < 1; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%07d", x), fmt.Sprintf("val-%07d", x))
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkGolangMapThroughputBatch_1(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		for j := 0; j < 1; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%07d", x)] = fmt.Sprintf("val-%07d", x)
				m.mu.Unlock()
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkConMapThroughputBatch_16(b *testing.B) {
	cm := NewConMap()
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(16)
		for j := 0; j < 16; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%07d", x), fmt.Sprintf("val-%07d", x))
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkGolangMapThroughputBatch_16(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(16)
		for j := 0; j < 16; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%07d", x)] = fmt.Sprintf("val-%07d", x)
				m.mu.Unlock()
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkConMapThroughputBatch_32(b *testing.B) {
	cm := NewConMap()
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(32)
		for j := 0; j < 32; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%07d", x), fmt.Sprintf("val-%07d", x))
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkGolangMapThroughputBatch_32(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(32)
		for j := 0; j < 32; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%07d", x)] = fmt.Sprintf("val-%07d", x)
				m.mu.Unlock()
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkConMapThroughputBatch_64(b *testing.B) {
	cm := NewConMap()
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(64)
		for j := 0; j < 64; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%07d", x), fmt.Sprintf("val-%07d", x))
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkGolangMapThroughputBatch_64(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(64)
		for j := 0; j < 64; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%07d", x)] = fmt.Sprintf("val-%07d", x)
				m.mu.Unlock()
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkConMapThroughputBatch_128(b *testing.B) {
	cm := NewConMap()
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(128)
		for j := 0; j < 128; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%07d", x), fmt.Sprintf("val-%07d", x))
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkGolangMapThroughputBatch_128(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	wg := new(sync.WaitGroup)
	for i := 0; i < b.N; i++ {
		wg.Add(128)
		for j := 0; j < 128; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%07d", x)] = fmt.Sprintf("val-%07d", x)
				m.mu.Unlock()
			}(j)
		}
		wg.Wait()
	}
}
