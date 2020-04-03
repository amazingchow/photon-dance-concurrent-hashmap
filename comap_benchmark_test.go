package comap

import (
	"fmt"
	"sync"
	"testing"
)

type TestUsageMap struct {
	mu sync.Mutex
	m  map[string]string
}

func BenchmarkCoMapThroughputBatch1(b *testing.B) {
	cm := NewCoMap()
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		for j := 0; j < 1; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%d", x), fmt.Sprintf("val-%d", x))
			}(j)
		}
	}
}

func BenchmarkGolangMapThroughputBatch1(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		for j := 0; j < 1; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%d", x)] = fmt.Sprintf("val-%d", x)
				m.mu.Unlock()
			}(j)
		}
	}
}

func BenchmarkCoMapThroughputBatch16(b *testing.B) {
	cm := NewCoMap()
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(16)
		for j := 0; j < 16; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%d", x), fmt.Sprintf("val-%d", x))
			}(j)
		}
	}
}

func BenchmarkGolangMapThroughputBatch16(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(16)
		for j := 0; j < 16; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%d", x)] = fmt.Sprintf("val-%d", x)
				m.mu.Unlock()
			}(j)
		}
	}
}

func BenchmarkCoMapThroughputBatch32(b *testing.B) {
	cm := NewCoMap()
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(32)
		for j := 0; j < 32; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%d", x), fmt.Sprintf("val-%d", x))
			}(j)
		}
	}
}

func BenchmarkGolangMapThroughputBatch32(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(32)
		for j := 0; j < 32; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%d", x)] = fmt.Sprintf("val-%d", x)
				m.mu.Unlock()
			}(j)
		}
	}
}

func BenchmarkCoMapThroughputBatch64(b *testing.B) {
	cm := NewCoMap()
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(64)
		for j := 0; j < 64; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%d", x), fmt.Sprintf("val-%d", x))
			}(j)
		}
	}
}

func BenchmarkGolangMapThroughputBatch64(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(64)
		for j := 0; j < 64; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%d", x)] = fmt.Sprintf("val-%d", x)
				m.mu.Unlock()
			}(j)
		}
	}
}

func BenchmarkCoMapThroughputBatch128(b *testing.B) {
	cm := NewCoMap()
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(128)
		for j := 0; j < 128; j++ {
			go func(x int) {
				defer wg.Done()
				cm.Store(fmt.Sprintf("key-%d", x), fmt.Sprintf("val-%d", x))
			}(j)
		}
	}
}

func BenchmarkGolangMapThroughputBatch128(b *testing.B) {
	m := &TestUsageMap{m: make(map[string]string)}
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(128)
		for j := 0; j < 128; j++ {
			go func(x int) {
				defer wg.Done()
				m.mu.Lock()
				m.m[fmt.Sprintf("key-%d", x)] = fmt.Sprintf("val-%d", x)
				m.mu.Unlock()
			}(j)
		}
	}
}
