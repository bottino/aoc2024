package day21

import (
	"math"
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
		{'7', '7'}: {"A"},
		{'8', '8'}: {"A"},
		{'4', '4'}: {"A"},
		{'5', '5'}: {"A"},
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

func TestCodeAllLevels(t *testing.T) {
	numPad := getShortestPaths(numKeys)
	arrowPad := getShortestPaths(arrowKeys)
	code := "029A"

	expSeq := "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"

	seqs := process(code, []Pad{numPad, arrowPad, arrowPad})

	if !slices.Contains(seqs, expSeq) {
		t.Errorf("Couldn't find sequence in %v", seqs)
	}

	minL := math.MaxInt
	maxL := 0
	for _, s := range seqs {
		if len(s) > maxL {
			maxL = len(s)
		}
		if len(s) < minL {
			minL = len(s)
		}
	}

	if minL != len(expSeq) || maxL != len(expSeq) {
		t.Errorf("Sequences don't all have the same length %d, %d", minL, maxL)
	}
}
