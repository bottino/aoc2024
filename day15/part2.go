package day15

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/bottino/aoc2024/vec"
)

func Part2(input string) any {
	sys := readInputPart2(input)
	sys.run()
	return sys.getCoordinates()
}

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

func (s *System2) run() {
	for s.numSteps < len(s.instructions) {
		s.step()
	}
}

func (s *System2) getCoordinates() int {
	var coordinates int
	for _, b := range s.boxes {
		minX := min(b.left.X, b.right.X)
		minY := min(b.left.Y, b.right.Y)
		coordinates += 100*minX + minY
	}

	// the boxes all appear twice in the map
	return coordinates / 2
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

func (s *System2) step() {
	instr := s.instructions[s.numSteps]
	newPos := s.robot.Add(instr)
	// do nothing if we hit a wall
	if s.walls[newPos] {
		s.numSteps++
		return
	}

	var boxes []*Box
	if b, ok := s.boxes[newPos]; ok {
		canMove := s.canMoveBox(b, instr, &boxes)
		if canMove {
			s.robot = newPos
			s.moveBoxes(boxes, instr)
		}
	} else {
		s.robot = newPos
	}

	// print
	// fmt.Printf("Step %d/%d: %v\n", s.numSteps, len(s.instructions), instr)
	// s.printWithPause(50 * time.Millisecond)

	s.numSteps++
}

func (s *System2) printWithPause(pause time.Duration) {
	fmt.Println(s.print())
	time.Sleep(pause)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (s *System2) canMoveBox(box *Box, dir vec.Coord, toMove *[]*Box) bool {
	*toMove = append(*toMove, box)
	canMove := true
	for _, pos := range []vec.Coord{box.left, box.right} {
		newPos := pos.Add(dir)
		if s.walls[newPos] {
			return false
		}

		b, ok := s.boxes[newPos]
		if ok && *b != *box {
			canMove = canMove && s.canMoveBox(b, dir, toMove)
		}
	}

	return canMove
}

func (s *System2) moveBoxes(boxes []*Box, dir vec.Coord) {
	for _, box := range boxes {
		delete(s.boxes, box.left)
		delete(s.boxes, box.right)
	}
	for _, box := range boxes {
		newLeft, newRight := box.left.Add(dir), box.right.Add(dir)
		newBox := Box{newLeft, newRight}
		s.boxes[newLeft] = &newBox
		s.boxes[newRight] = &newBox
	}
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
