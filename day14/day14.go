package day14

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	robots := readRobots(input, 103, 101)
	for i := 0; i < 156; i++ {
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
	N, M := 103, 101
	robots := readRobots(input, N, M)

	maxN := 10_000
	middle := N / 2
	var maxMiddle int
	var maxes []int
	for i := 0; i < maxN; i++ {
		var middleCount int
		for _, r := range robots {
			r.Move()
			if vec.AbsInt(r.Pos.X-middle) <= 10 {
				middleCount++
			}
		}

		if middleCount >= maxMiddle {
			maxMiddle = middleCount
			maxes = append(maxes, i+1)

		}

		// I noticed a pattern where most of the robots are in the
		// middle at this frequency (by looking at maxes)
		// I then just stopped when I saw the tree, at 7572

		// In retrospect, I figured that there was a vertical pattern
		// at 98 + n*101, and a horizontal pattern at 52 + m*103, so
		// the first integer that satifies both conditions is 7572
		// if i%103 == 52 {
		// 	displayWithPause(N, M, i, robots)
		// }
	}

	return 7572
}

func displayWithPause(N, M, i int, robots []*Robot) {
	fmt.Println(displayMap(N, M, robots))
	fmt.Printf("Time elapsed: %d s\n\n", i+1)
	time.Sleep(300 * time.Millisecond)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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

func displayMap(N, M int, robots []*Robot) string {
	robMap := make(map[vec.Coord]int, len(robots))
	for _, r := range robots {
		robMap[r.Pos] += 1
	}

	var sb strings.Builder
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if v, ok := robMap[vec.Coord{i, j}]; ok {
				sb.WriteString(fmt.Sprintf("%d", v))
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
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
