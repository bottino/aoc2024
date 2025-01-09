package day15

import (
	"fmt"
	"strings"

	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	sys := readInput(input)
	// fmt.Println(sys.print())
	sys.run()
	return sys.getCoordinate()
}

type System struct {
	robot        vec.Coord
	obstacles    map[vec.Coord]rune
	instructions []vec.Coord
	numSteps     int
	N, M         int
}

func (s *System) run() {
	for s.numSteps < len(s.instructions) {
		s.step()
	}
}

func (s *System) step() {
	instr := s.instructions[s.numSteps]
	newPos := s.robot.Add(instr)
	if v, ok := s.obstacles[newPos]; ok {
		switch v {
		case '#': // do nothing if we hit a wall
		case 'O':
			canMove := s.moveBox(newPos, instr)
			if canMove {
				s.robot = newPos
			}
		}
	} else {
		s.robot = newPos
	}

	// print
	// fmt.Printf("Step %d: %v\n", s.numSteps, instr)
	// fmt.Println(s.print())

	s.numSteps++
}

func (s *System) moveBox(pos vec.Coord, dir vec.Coord) bool {
	newPos := pos.Add(dir)
	v, ok := s.obstacles[newPos]
	if !ok {
		delete(s.obstacles, pos)
		s.obstacles[newPos] = 'O'
		return true
	}
	switch v {
	case '#':
		return false
	case 'O':
		canMove := s.moveBox(newPos, dir)
		if canMove {
			delete(s.obstacles, pos)
			s.obstacles[newPos] = 'O'
		}
		return canMove
	default:
		fmt.Printf("Unexpected value %v at pos %v\n", v, newPos)
		return false
	}
}

func (s *System) getCoordinate() int {
	var coordinates int
	for coord, obs := range s.obstacles {
		if obs != 'O' {
			continue
		}

		coordinates += 100*coord.X + coord.Y
	}
	return coordinates
}

func (s *System) print() string {
	var sb strings.Builder
	for i := 0; i < s.N; i++ {
		for j := 0; j < s.M; j++ {
			if (vec.Coord{i, j} == s.robot) {
				sb.WriteRune('@')
				continue
			}
			if v, ok := s.obstacles[vec.Coord{i, j}]; ok {
				sb.WriteRune(v)
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func readInput(input string) System {
	sys := System{obstacles: make(map[vec.Coord]rune)}
	parts := strings.Split(input, "\n\n")
	for i, line := range strings.Split(parts[0], "\n") {
		for j, char := range line {
			switch char {
			case '#':
				sys.obstacles[vec.Coord{i, j}] = '#'
			case 'O':
				sys.obstacles[vec.Coord{i, j}] = 'O'
			case '.':
				continue
			case '@':
				sys.robot = vec.Coord{X: i, Y: j}
			default:
				fmt.Printf("Unknown char %v\n", char)
			}
			sys.N, sys.M = i+1, j+1
		}
	}

	for _, char := range strings.Replace(parts[1], "\n", "", -1) {
		switch char {
		case '>':
			sys.instructions = append(sys.instructions, vec.Right)
		case '<':
			sys.instructions = append(sys.instructions, vec.Left)
		case '^':
			sys.instructions = append(sys.instructions, vec.Up)
		case 'v':
			sys.instructions = append(sys.instructions, vec.Down)
		}
	}

	return sys
}
