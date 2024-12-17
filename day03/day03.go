package day03

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

func Part1(input string) any {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	var sum int
	for _, m := range matches {
		sum += toInt(m[1]) * toInt(m[2])
	}

	return sum
}

func Part2(input string) any {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	enabled := true
	var sum int
	for _, m := range matches {
		switch m[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				sum += toInt(m[1]) * toInt(m[2])
			}
		}
	}

	return sum
}

func toInt(x string) int {
	xInt, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("String conversion failed for %s\n", x)
	}

	return xInt
}
