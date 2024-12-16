package day16

import (
	"testing"
)

func TestPQueue(t *testing.T) {
	fruits := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := NewPQueue[string]()

	for k, v := range fruits {
		pq.AddWithRank(k, v)
	}

	expected := []string{"apple", "banana", "pear"}
	for _, exp := range expected {
		popped := pq.PopMin()
		if popped != exp {
			t.Errorf("Expected %s, got %s", exp, popped)
		}
	}
}
