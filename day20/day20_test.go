package day20

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var input = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`

func TestReadTrack(t *testing.T) {
	track, distances := readInput(input)
	if len(track)-1 != 84 {
		t.Errorf("Wrong track length %d", len(track))
	}
	if len(distances)-1 != 84 {
		t.Errorf("Wrong distances length %d", len(track))
	}
}

func TestCheats(t *testing.T) {
	cheats := findCheats(input, 20, 50)

	cheatMap := make(map[int]int)
	for _, c := range cheats {
		v := cheatMap[c]
		cheatMap[c] = v + 1
	}

	expCheat := map[int]int{
		50: 32,
		52: 31,
		54: 29,
		56: 39,
		58: 25,
		60: 23,
		62: 20,
		64: 19,
		66: 12,
		68: 14,
		70: 12,
		72: 22,
		74: 4,
		76: 3,
	}

	if diff := cmp.Diff(expCheat, cheatMap); diff != "" {
		t.Error(diff)
	}
}
