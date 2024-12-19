package day19

import (
	"strings"
)

func Part1(input string) any {
	patterns, designs := readInputFile(input)

	var numPossible int
	memo := make(map[string]bool)
	for _, d := range designs {
		if isPossible(d, patterns, memo) {
			numPossible++
		}
	}

	return numPossible
}

func Part2(input string) any {
	patterns, designs := readInputFile(input)

	var allWays int
	memo := make(map[string]int)
	for _, d := range designs {
		allWays += getWays(d, patterns, memo)
	}

	return allWays
}

func isPossible(design string, patterns []string, memo map[string]bool) bool {
	if v, ok := memo[design]; ok {
		return v
	}

	for _, p := range patterns {
		nP := len(p)
		nD := len(design)
		if nD >= nP && len(design[:nP]) == nP && design[:nP] == p {
			if nD == nP {
				memo[design] = true
				return true
			}
			rem := design[nP:]

			if isPossible(rem, patterns, memo) {
				memo[design] = true
				return true
			}
		}
	}

	memo[design] = false
	return false
}

func getWays(design string, patterns []string, memo map[string]int) int {
	if v, ok := memo[design]; ok {
		return v
	}

	var ways int
	for _, p := range patterns {
		nP := len(p)
		nD := len(design)
		if nD >= nP && len(design[:nP]) == nP && design[:nP] == p {
			if nD == nP {
				ways++
			}
			rem := design[nP:]

			ways += getWays(rem, patterns, memo)
		}
	}

	memo[design] = ways
	return ways
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
