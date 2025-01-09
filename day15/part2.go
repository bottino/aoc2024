package day15

import (
	"fmt"
	"strings"

	"github.com/bottino/aoc2024/vec"
)

type System2 struct {
	robot        vec.Coord
	walls        map[vec.Coord]bool
	boxes        map[vec.Coord]*Box
	instructions []vec.Coord
	numSteps     int
	N, M         int
}

type Box struct {
	left, right vec.Coord
}

func (s *System2) print() string {
	var sb strings.Builder
	for i := 0; i < s.N; i++ {
		for j := 0; j < s.M; j++ {
			u := vec.Coord{i, j}
			if u == s.robot {
				sb.WriteRune('@')
				continue
			}
			if s.walls[u] {
				sb.WriteRune('#')
				continue
			}
			if b, ok := s.boxes[u]; ok {
				if b.left == u {
					sb.WriteRune('[')
				} else if b.right == u {
					sb.WriteRune(']')
				} else {
					fmt.Printf("ERROR: Box doesn't contain coord: %v, %v", *b, u)
				}
				continue
			}

			sb.WriteRune('.')
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func readInputPart2(input string) System2 {
	sys := System2{walls: make(map[vec.Coord]bool), boxes: make(map[vec.Coord]*Box)}
	parts := strings.Split(input, "\n\n")
	for i, line := range strings.Split(parts[0], "\n") {
		for j := 0; j < 2*len(line); j += 2 {
			switch line[j/2] {
			case '#':
				sys.walls[vec.Coord{i, j}] = true
				sys.walls[vec.Coord{i, j + 1}] = true
			case 'O':
				box := Box{left: vec.Coord{i, j}, right: vec.Coord{i, j + 1}}
				sys.boxes[vec.Coord{i, j}] = &box
				sys.boxes[vec.Coord{i, j + 1}] = &box
			case '.':
				continue
			case '@':
				sys.robot = vec.Coord{X: i, Y: j}
			default:
				fmt.Printf("Unknown char %v\n", line[j/2])
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
