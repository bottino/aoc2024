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

	robots := readRobots(input, 7, 11)
	expected := []*Robot{
		{vec.Coord{4, 0}, vec.Coord{-3, 3}, 7, 11},
		{vec.Coord{3, 6}, vec.Coord{-3, -1}, 7, 11},
		{vec.Coord{3, 10}, vec.Coord{2, -1}, 7, 11},
	}

	if diff := cmp.Diff(expected, robots); diff != "" {
		t.Errorf("Unexpected diff in robots: %s", diff)
	}
}

func TestMoveRobot(t *testing.T) {
	r := Robot{vec.Coord{4, 2}, vec.Coord{-3, 2}, 7, 11}
	exp := []vec.Coord{
		{1, 4},
		{5, 6},
		{2, 8},
		{6, 10},
		{3, 1},
	}

	for i := 0; i < 5; i++ {
		r.Move()
		if diff := cmp.Diff(exp[i], r.Pos); diff != "" {
			t.Errorf("Wrong pos at time %d: %s", i+1, diff)
		}
	}
}
