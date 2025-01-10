package day25

import (
	"strings"
)

func Part1(input string) any {
	keys, locks := readInput(input)

	var fits int
	for _, k := range keys {
		for _, l := range locks {
			var pinFits int
			for i := 0; i < 5; i++ {
				if k[i]+l[i] <= 5 {
					pinFits++
				}
			}
			if pinFits == 5 {
				fits++
			}
		}
	}

	return fits
}

func Part2(input string) any {
	return "Done!"
}

func readInput(input string) (keys []pins, locks []pins) {
	for _, pattern := range strings.Split(input, "\n\n") {
		heights := pins{-1, -1, -1, -1, -1}
		var isKey bool
		for i, line := range strings.Split(pattern, "\n") {
			if i == 0 && line == "#####" {
				isKey = true
			}

			for j, char := range line {
				if char == '#' {
					heights[j]++
				}
			}
		}

		if isKey {
			keys = append(keys, heights)
		} else {
			locks = append(locks, heights)
		}
	}

	return keys, locks
}

type pins [5]int
