package day19

import "testing"

func TestIsPossible(t *testing.T) {
	patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "bR"}

	if !isPossible("bwurrg", patterns) {
		t.Error("Should work")
	}
}
