package day14

import (
	"testing"

	"github.com/bottino/aoc2024/vec"
	"github.com/google/go-cmp/cmp"
)

func TestReadRobots(t *testing.T) {
	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2`

	robots := readRobots(input)
	expected := []Robot{
		{vec.Coord{4, 0}, vec.Coord{-3, 3}},
		{vec.Coord{3, 6}, vec.Coord{-3, -1}},
		{vec.Coord{3, 10}, vec.Coord{2, -1}},
	}

	if diff := cmp.Diff(expected, robots); diff != "" {
		t.Errorf("Unexpected diff in robots: %s", diff)
	}
}
