package betterguid

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

const (
	// for manual debugging
	showIds = false
)

func Test(t *testing.T) {
	id1 := New()
	if len(id1) != 20 {
		t.Fatalf("len(id1) != 20 (=%d)", len(id1))
	}
	id2 := New()
	if len(id2) != 20 {
		t.Fatalf("len(id1) != 20 (=%d)", len(id2))
	}
	if id1 == id2 {
		t.Fatalf("generated same ids (id1: '%s', id2: '%s')", id1, id2)
	}
	if showIds {
		fmt.Printf("%s\n", id1)
		fmt.Printf("%s\n", id2)
	}
}

func doMany(t *testing.T, wg *sync.WaitGroup) {
	ids := make(map[string]bool)
	prev := ""
	for i := 0; i < 1000000; i++ {
		id := New()
		if _, exists := ids[id]; exists {
			t.Fatalf("generated duplicate id '%s'", id)
		}
		ids[id] = true
		if prev != "" {
			if id <= prev {
				t.Fatalf("id ('%s') must be > prev ('%s')", id, prev)
			}
		}
		prev = id
	}
	wg.Done()
}

func TestMany(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go doMany(t, &wg)
	}
	wg.Wait()
}
