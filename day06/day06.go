package day06

import (
	"fmt"
	"strings"
)

func Part1(input string) (solution int) {
	grid, guard := readInput(input)
	solution++ // the first square is visited
	for {
		pos := Coord{guard.Pos[0] + guard.Dir[0], guard.Pos[1] + guard.Dir[1]}
		newSquare, ok := grid[pos]
		if ok == false {
			fmt.Println("Guard exited at", pos)
			break
		}

		switch newSquare {
		case '.':
			grid[pos] = 'X'
			guard.Pos = pos
			solution++
		case 'X':
			guard.Pos = pos
		case '#':
			guard.TurnRight()
		}
	}

	return
}

func Part2(input string) (solution int) {
	fmt.Println("No solution yet for day 6, part 2")
	return
}

type Coord [2]int

var (
	up    = Coord{-1, 0}
	down  = Coord{1, 0}
	left  = Coord{0, -1}
	right = Coord{0, 1}
)

type Guard struct {
	Pos Coord
	Dir Coord
}

func (g *Guard) TurnRight() {
	switch g.Dir {
	case up:
		g.Dir = right
	case right:
		g.Dir = down
	case down:
		g.Dir = left
	case left:
		g.Dir = up
	}
}

type Grid map[Coord]rune

func readInput(input string) (grid Grid, guard Guard) {
	grid = make(Grid)
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			square := Coord{i, j}
			switch char {
			case '^':
				guard = Guard{square, up}
				grid[square] = 'X'
			case '>':
				guard = Guard{square, right}
				grid[square] = 'X'
			case '<':
				guard = Guard{square, left}
				grid[square] = 'X'
			case 'v':
				guard = Guard{square, down}
				grid[square] = 'X'
			default:
				grid[square] = char
			}
		}
	}

	return
}
