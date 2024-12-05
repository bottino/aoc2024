package day02

import (
	_ "embed"
	"strconv"
	"strings"
)

func Part1(input string) int {
	reports := readLevels(input)
	var numSafe int
	for i := range reports {
		if isSafe(reports[i]) {
			numSafe++
		}
	}
	return numSafe
}

func Part2(input string) int {
	reports := readLevels(input)
	var numSafe int
	for i := range reports {
		if isSafeLenient(reports[i]) {
			numSafe++
		}
	}
	return numSafe
}

func isSafeLenient(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := range levels {
		var levelsOneRemoved []int
		levelsOneRemoved = append(levelsOneRemoved, levels[:i]...)
		levelsOneRemoved = append(levelsOneRemoved, levels[i+1:]...)
		if isSafe(levelsOneRemoved) {
			return true
		}
	}

	return false
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
