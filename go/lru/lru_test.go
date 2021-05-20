package lru

import (
	"testing"
)

func TestLRU_Put(t *testing.T) {
	lru := NewLRU(4)
	lru.Put(1, "a")
	lru.Put(1, "a")
	lru.Put(2, "b")
	lru.Put(3, "c")
	lru.Get(1)
	lru.Put(0, "d")
	lru.Put(1, "F")

	if lru.Get(1) != "F" {
		t.Error("fail")
	}

	lru.Print()
}

