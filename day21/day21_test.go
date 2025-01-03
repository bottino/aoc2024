package day21

import (
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

	memo := make(map[string]string)
	seq := processCode("A"+code, numPad, &memo)

	if len(seq) != len(expSeq) {
		t.Errorf("Expected length %d, got %d", len(expSeq), len(seq))
	}
}

func TestCodeAllLevels(t *testing.T) {
	numPad := getShortestPaths(numKeys)
	arrowPad := getShortestPaths(arrowKeys)
	code := "029A"

	expSeq := "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"

	seq := process(code, []Pad{numPad, arrowPad, arrowPad})

	if len(seq) != len(expSeq) {
		t.Errorf("Expected length %d, got %d", len(expSeq), len(seq))
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
