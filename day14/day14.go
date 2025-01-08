package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	fmt.Println("No solution yet for day 14, part 1")
	return 0
}

func Part2(input string) any {
	fmt.Println("No solution yet for day 14, part 2")
	return 0
}

type Robot struct {
	Pos vec.Coord
	Dir vec.Coord
}

func readRobots(input string) (robots []Robot) {
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
		r := Robot{vec.Coord{nums[1], nums[0]}, vec.Coord{nums[3], nums[2]}}
		robots = append(robots, r)
	}

	return robots
}
