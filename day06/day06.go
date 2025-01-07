package day06

import (
	"strings"
	"sync"

	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	grid, guard := readInput(input)
	sys := System{guard.Clone(), grid, vec.Coord{X: -1, Y: -1}}

	visited, _ := runSystem(sys)

	return len(visited) + 1 // the starting square
}

func Part2(input string) any {
	grid, guard := readInput(input)

	sys := System{guard.Clone(), grid, vec.Coord{X: -1, Y: -1}}
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

func runSystem(sys System) (visited map[vec.Coord]bool, isLoop bool) {
	visited = make(map[vec.Coord]bool)
	obsHits := make(map[Hit]bool)
	for {
		pos := sys.Guard.Pos.Add(sys.Guard.Dir)
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
			hit := Hit{pos.X, pos.Y, sys.Guard.Dir.X, sys.Guard.Dir.Y}
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
	ExtraObstacle vec.Coord
}

type Hit [4]int

type Guard struct {
	Pos vec.Coord
	Dir vec.Coord
}

func (g *Guard) Clone() Guard {
	return Guard{Pos: g.Pos, Dir: g.Dir}
}

func (g *Guard) TurnRight() {
	switch g.Dir {
	case vec.Up:
		g.Dir = vec.Right
	case vec.Right:
		g.Dir = vec.Down
	case vec.Down:
		g.Dir = vec.Left
	case vec.Left:
		g.Dir = vec.Up
	}
}

type Grid map[vec.Coord]rune

func readInput(input string) (grid Grid, guard Guard) {
	grid = make(Grid)
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, char := range line {
			square := vec.Coord{X: i, Y: j}
			switch char {
			case '^':
				guard = Guard{square, vec.Up}
				grid[square] = 'X'
			case '>':
				guard = Guard{square, vec.Right}
				grid[square] = 'X'
			case '<':
				guard = Guard{square, vec.Left}
				grid[square] = 'X'
			case 'v':
				guard = Guard{square, vec.Down}
				grid[square] = 'X'
			default:
				grid[square] = char
			}
		}
	}

	return
}
