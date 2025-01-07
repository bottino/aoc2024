package graphs

import (
	"testing"
)

func TestSetAddingAndRemoving(t *testing.T) {
	s1 := NewSet("a", "b", "c", "a")
	s1.Remove("b")
	s2 := NewSet("c", "a")

	if !s1.Equal(s2) {
		t.Errorf("Sets are not equal: %s %s", s1, s2)
	}
}

func TestUnionIntersection(t *testing.T) {
	s1 := NewSet("a", "b", "c")
	s2 := NewSet("c", "e", "f")
	if !((s1.Union(s2)).Equal(NewSet("b", "c", "a", "e", "f"))) {
		t.Errorf("Failed union: %s %s", s1, s2)
	}

	if !((s1.Intersection(s2)).Equal(NewSet("c"))) {
		t.Errorf("Failed intersection: %s %s", s1, s2)
	}
}
