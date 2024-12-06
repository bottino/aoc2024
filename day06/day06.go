package day06

import (
	"fmt"
	"strings"
)

func Part1(input string) (solution int) {
	grid, guard, _, _ := readInput(input)
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
	grid, guard, nx, ny := readInput(input)
	for i := range nx {
		for j := range ny {
			oPos := Coord{i, j}
			if grid[oPos] != '.' {
				continue
			}

			sys := System{
				Guard:    guard.Clone(),
				Grid:     grid,
				Obstacle: oPos,
				ObsHits:  make(map[Hit]bool),
			}

			if isSystemLoop(sys) {
				//Found loop with obstacle at oPos
				solution++
			}
		}
	}
	return
}

func isSystemLoop(sys System) bool {
	for {
		pos := Coord{sys.Guard.Pos[0] + sys.Guard.Dir[0], sys.Guard.Pos[1] + sys.Guard.Dir[1]}
		newSquare, ok := sys.Grid[pos]
		if ok == false {
			// Guard exists through the edge
			return false
		}

		// Check if we hit the extra obstacle
		if pos == sys.Obstacle {
			newSquare = '#'
		}

		switch newSquare {
		case '.':
			sys.Guard.Pos = pos
		case 'X':
			sys.Guard.Pos = pos
		case '#':
			// If we've hit the same obstacle facing the same direction, we're in a loop
			hit := Hit{pos[0], pos[1], sys.Guard.Dir[0], sys.Guard.Dir[1]}
			if _, ok := sys.ObsHits[hit]; ok {
				// There's a loop
				return true
			}
			sys.ObsHits[hit] = true
			sys.Guard.TurnRight()
		}
	}
}

type System struct {
	Guard    Guard
	Grid     Grid
	Obstacle Coord
	ObsHits  map[Hit]bool
}

type Hit [4]int
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

func (g *Guard) Clone() Guard {
	return Guard{Pos: g.Pos, Dir: g.Dir}
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

func readInput(input string) (grid Grid, guard Guard, nx int, ny int) {
	grid = make(Grid)
	lines := strings.Split(input, "\n")
	nx = len(lines)
	for i, line := range lines {
		ny = len(line)
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
