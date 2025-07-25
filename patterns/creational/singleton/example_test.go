package singleton

import (
	"sync"
	"testing"
)

func TestSingletonInstance(t *testing.T) {
	var wg sync.WaitGroup
	instances := make([]*singleton, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			instances[idx] = GetInstance()
		}(i)
	}
	wg.Wait()

	for i := 1; i < 10; i++ {
		if instances[i] != instances[0] {
			t.Errorf("Expected the same instance, got different ones")
		}
	}
}