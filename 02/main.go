package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func init() {
	example = strings.TrimRight(example, "\n")
	input = strings.TrimRight(input, "\n")
}

func main() {
	var part int
	flag.IntVar(&part, "p", 1, "The part of the puzzle")
	var useExample bool
	flag.BoolVar(&useExample, "e", false, "Use the example as input")
	flag.Parse()

	if useExample {
		input = example
	}

	var solution int
	if part == 1 {
		solution = part1(input)
	} else {
		solution = part2(input)
	}

	fmt.Println(solution)
}

func part1(input string) int {
	reports := readLevels(input)
	var numSafe int
	for i := range reports {
		if isSafe(reports[i]) {
			numSafe++
		}
	}
	return numSafe
}

func part2(input string) int {
	// part 2 here
	return 0
}

func isSafe(levels []int) bool {
	var isIncreasing bool
	for i := 1; i < len(levels); i++ {
		var diff int = levels[i] - levels[i-1]

		if i == 1 {
			isIncreasing = diff > 0
		}

		if isIncreasing != (diff > 0) {
			return false
		}

		if absInt(diff) < 1 || absInt(diff) > 3 {
			return false
		}
	}

	return true
}

func absInt(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func readLevels(input string) (reports [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var levels []int
		for _, levelStr := range strings.Split(line, " ") {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				panic("Error converting to int")
			}

			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}

	return reports
}
