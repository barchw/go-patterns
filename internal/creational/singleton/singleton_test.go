package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance_SameInstance(t *testing.T) {
	a := GetInstance()
	b := GetInstance()

	if a != b {
		t.Fatal("GetInstance returned different instances")
	}

	a.Set("key", "value")
	if got := b.Get("key"); got != "value" {
		t.Errorf("Get(key) = %q, want %q", got, "value")
	}
}

func TestGetInstance_ConcurrentAccess(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	instances := make([]*ConfigManager, goroutines)

	wg.Add(goroutines)
	for i := range goroutines {
		go func(idx int) {
			defer wg.Done()
			instances[idx] = GetInstance()
		}(i)
	}
	wg.Wait()

	first := instances[0]
	for i, inst := range instances {
		if inst != first {
			t.Fatalf("instance[%d] differs from instance[0]", i)
		}
	}
}
