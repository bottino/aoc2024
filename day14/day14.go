package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	robots := readRobots(input, 103, 101)
	for i := 0; i < 100; i++ {
		for _, r := range robots {
			r.Move()
		}
	}

	quad := map[int]int{0: 0, 1: 0, 2: 0, 3: 0}
	for _, r := range robots {
		if q := r.Quadrant(); q >= 0 {
			quad[r.Quadrant()] += 1
		}
	}

	safetyFactor := 1
	for _, v := range quad {
		safetyFactor *= v
	}

	return safetyFactor
}

func Part2(input string) any {
	fmt.Println("No solution yet for day 14, part 2")
	return 0
}

type Robot struct {
	Pos  vec.Coord
	Dir  vec.Coord
	N, M int
}

func (r *Robot) Move() {
	p := r.Pos.Add(r.Dir)
	if p.X >= r.N {
		p.X -= r.N
	}
	if p.X < 0 {
		p.X += r.N
	}
	if p.Y >= r.M {
		p.Y -= r.M
	}
	if p.Y < 0 {
		p.Y += r.M
	}
	r.Pos = p
}

func (r *Robot) Quadrant() int {
	switch {
	case r.Pos.X < r.N/2 && r.Pos.Y < r.M/2:
		return 0
	case r.Pos.X > r.N/2 && r.Pos.Y < r.M/2:
		return 1
	case r.Pos.X < r.N/2 && r.Pos.Y > r.M/2:
		return 2
	case r.Pos.X > r.N/2 && r.Pos.Y > r.M/2:
		return 3
	default:
		return -1
	}
}

func readRobots(input string, N, M int) (robots []*Robot) {
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range strings.Split(input, "\n") {
		matches := re.FindAllString(line, -1)
		var nums []int
		for _, m := range matches {
			n, err := strconv.Atoi(m)
			if err != nil {
				fmt.Printf("Could not convert %s to int", m)
			}
			nums = append(nums, n)
		}

		// Reverse x and y because of the convention of the vec package
		r := Robot{
			vec.Coord{nums[1], nums[0]},
			vec.Coord{nums[3], nums[2]},
			N, M,
		}
		robots = append(robots, &r)
	}

	return robots
}
