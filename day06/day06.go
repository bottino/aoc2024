package day06

import (
	"fmt"
	"strings"
)

func Part1(input string) (solution int) {
	grid, guard := readInput(input)
	solution++ // the first square is visited
	for {
		x, y := guard.X+guard.D[0], guard.Y+guard.D[1]
		newSquare, ok := grid[x][y]
		if ok == false {
			fmt.Println("Guard exited at", x, y)
			break
		}

		switch newSquare {
		case '.':
			grid[x][y] = 'X'
			guard.X, guard.Y = x, y
			solution++
		case 'X':
			guard.X, guard.Y = x, y
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

type Dir [2]int

var (
	up    = Dir{-1, 0}
	down  = Dir{1, 0}
	left  = Dir{0, -1}
	right = Dir{0, 1}
)

type Guard struct {
	X, Y int
	D    Dir
}

func (g *Guard) TurnRight() {
	switch g.D {
	case up:
		g.D = right
	case right:
		g.D = down
	case down:
		g.D = left
	case left:
		g.D = up
	}
}

type Grid map[int]map[int]rune

func readInput(input string) (grid Grid, guard Guard) {
	grid = make(Grid)
	for i, line := range strings.Split(input, "\n") {
		grid[i] = make(map[int]rune)
		for j, char := range line {
			switch char {
			case '^':
				guard = Guard{i, j, up}
				grid[i][j] = 'X'
			case '>':
				guard = Guard{i, j, right}
				grid[i][j] = 'X'
			case '<':
				guard = Guard{i, j, left}
				grid[i][j] = 'X'
			case 'v':
				guard = Guard{i, j, down}
				grid[i][j] = 'X'
			default:
				grid[i][j] = char
			}
		}
	}

	return
}
