package day21

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShortestPaths(t *testing.T) {
	keys := map[Coord]rune{
		{0, 0}: '7',
		{1, 0}: '8',
		{0, 1}: '4',
		{1, 1}: '5',
	}

	expected := Pad{
		{'7', '8'}: {">A"},
		{'8', '7'}: {"<A"},
		{'4', '5'}: {">A"},
		{'5', '4'}: {"<A"},
		{'7', '4'}: {"vA"},
		{'4', '7'}: {"^A"},
		{'8', '5'}: {"vA"},
		{'5', '8'}: {"^A"},
		{'7', '5'}: {">vA", "v>A"},
		{'5', '7'}: {"<^A", "^<A"},
		{'4', '8'}: {">^A", "^>A"},
		{'8', '4'}: {"<vA", "v<A"},
	}

	paths := getShortestPaths(keys)
	if diff := cmp.Diff(expected, paths); diff != "" {
		t.Error(diff)
	}
}

func TestCodeOneLevel(t *testing.T) {
	numPad := getShortestPaths(numKeys)
	code := "029A"

	expSeq := "<A^A>^^AvvvA"

	seqs := processCode(code, numPad)

	if !slices.Contains(seqs, expSeq) {
		t.Errorf("Couldn't find sequence in %v", seqs)
	}
}
