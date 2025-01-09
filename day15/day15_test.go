package day15

import (
	"fmt"
	"testing"
)

// We should not move boxes if another "thread" is blocking
func TestBackProp(t *testing.T) {
	testInput := `
#######
#...#.#
#...#.#
#..OO@#
#..O..#
#.....#
#######

<vv<<<^^
`

	exp := `
#############
##......##..#
##......##..#
##...[][]...#
##....[]....#
##....@.....#
#############
`

	sys := readInputPart2(testInput[1:])
	sys.run()
	actual := sys.print()
	if actual != exp[1:] {
		fmt.Println(exp)
		fmt.Println(actual)
		t.Error("Wrong end state (top expected, bottom actual)")
	}
}
