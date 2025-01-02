package graphs

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

	// Update with pear
	pq.AddWithRank("pear", 0)

	expected := []string{"pear", "apple", "banana"}
	for _, exp := range expected {
		popped := pq.PopMin()
		if popped != exp {
			t.Errorf("Expected %s, got %s", exp, popped)
		}
	}
}
