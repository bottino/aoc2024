package day06

import (
	"strings"
	"sync"
)

func Part1(input string) any {
	grid, guard := readInput(input)
	sys := System{guard.Clone(), grid, Coord{-1, -1}}

	visited, _ := runSystem(sys)

	return len(visited) + 1 // the starting square
}

func Part2(input string) any {
	grid, guard := readInput(input)

	sys := System{guard.Clone(), grid, Coord{-1, -1}}
	visited, _ := runSystem(sys)

	c := make(chan bool, len(visited))
	wg := sync.WaitGroup{}

	for k := range visited {
		// Don't process if existing obstacle or the starting position of the guard
		if grid[k] != '.' {
			continue
		}

		sys := System{guard.Clone(), grid, k}

		wg.Add(1)
		go findLoopInSystem(sys, c, &wg)
	}

	wg.Wait()
	close(c)

	var solution int
	for isLoop := range c {
		if isLoop {
			solution++
		}
	}

	return solution
}

func findLoopInSystem(sys System, c chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	_, isLoop := runSystem(sys)
	c <- isLoop
}

func runSystem(sys System) (visited map[Coord]bool, isLoop bool) {
	visited = make(map[Coord]bool)
	obsHits := make(map[Hit]bool)
	for {
		pos := Coord{sys.Guard.Pos[0] + sys.Guard.Dir[0], sys.Guard.Pos[1] + sys.Guard.Dir[1]}
		newSquare, ok := sys.Grid[pos]
		if ok == false {
			// Guard exists through the edge
			return visited, false
		}

		if pos == sys.ExtraObstacle {
			newSquare = '#'
		}

		switch newSquare {
		case '.':
			sys.Guard.Pos = pos
			visited[pos] = true
		case 'X':
			sys.Guard.Pos = pos
		case '#':
			// If we've hit the same obstacle facing the same direction, we're in a loop
			hit := Hit{pos[0], pos[1], sys.Guard.Dir[0], sys.Guard.Dir[1]}
			if _, ok := obsHits[hit]; ok {
				// There's a loop
				return visited, true
			}
			obsHits[hit] = true
			sys.Guard.TurnRight()
		}
	}
}

type System struct {
	Guard         Guard
	Grid          Grid
	ExtraObstacle Coord
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

func readInput(input string) (grid Grid, guard Guard) {
	grid = make(Grid)
	lines := strings.Split(input, "\n")
	for i, line := range lines {
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
