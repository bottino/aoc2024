package day20

import "testing"

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
	track := readInput(input)
	if len(track)-1 != 84 {
		t.Errorf("Wrong track length %d", len(track))
	}
}
