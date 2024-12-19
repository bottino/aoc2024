package day19

import (
	"fmt"
	"strings"
)

func Part1(input string) any {
	patterns, designs := readInputFile(input)

	var numPossible int
	for _, d := range designs {
		if isPossible(d, patterns) {
			numPossible++
		}
	}

	return numPossible
}

func Part2(input string) any {
	fmt.Println("No solution yet for day 19, part 2")
	return 0
}

func isPossible(design string, patterns []string) bool {
	for _, p := range patterns {
		nP := len(p)
		nD := len(design)
		if nD >= nP && len(design[:nP]) == nP && design[:nP] == p {
			if nD == nP {
				return true
			}
			rem := design[nP:]

			if isPossible(rem, patterns) {
				return true
			}
		}
	}

	return false
}

func readInput(input string) {
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
	}
}

func readInputFile(input string) ([]string, []string) {
	lines := strings.Split(input, "\n")

	// trim all strings in lines
	patterns := []string{}
	for _, pat := range strings.Split(lines[0], ",") {
		patterns = append(patterns, strings.TrimSpace(pat))
	}
	return patterns, lines[2:]
}
