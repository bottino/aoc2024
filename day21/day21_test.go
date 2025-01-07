package day21

import (
	"slices"
	"testing"

	"github.com/bottino/aoc2024/vec"
	"github.com/google/go-cmp/cmp"
)

func TestShortestPaths(t *testing.T) {
	keys := map[vec.Coord]rune{
		{X: 0, Y: 0}: '7',
		{X: 0, Y: 1}: '8',
		{X: 1, Y: 0}: '4',
		{X: 1, Y: 1}: '5',
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

func TestBuildSec(t *testing.T) {
	numPad := getShortestPaths(numKeys)
	code := "029A"

	expSeq := "<A^A>^^AvvvA"
	results := getSeqs(code, numPad)
	if !slices.Contains(results, expSeq) {
		t.Errorf("seq not in %v", results)
	}
}

func TestSplitCode(t *testing.T) {
	code := "A290A"
	left := code[:len(code)/2+1]
	right := code[len(code)/2:]
	if left != "A29" || right != "90A" {
		t.Errorf("Wrong split: %s %s", left, right)
	}
}
