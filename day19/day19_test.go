package day19

import "testing"

func TestIsPossible(t *testing.T) {
	patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "bR"}

	memo := make(map[string]bool)
	if !isPossible("bwurrg", patterns, memo) {
		t.Error("Should work")
	}
}
