package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	var sum int
	for _, match := range matches {
		lhs, err_l := strconv.Atoi(match[1])
		rhs, err_r := strconv.Atoi(match[2])
		if err_l != nil || err_r != nil {
			panic("String conversion failed")
		}

		sum += lhs * rhs
	}

	return sum
}

func part2(input string) int {
	dos := find_re_indices(input, `do\(\)`)
	donts := find_re_indices(input, `don't\(\)`)
	var sum int

	re_mul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	mul_matches := re_mul.FindAllStringSubmatchIndex(input, -1)
	for _, m := range mul_matches {
		idx := m[0]
		preceding_do := find_closest_preceding_index(idx, dos)
		preceding_dont := find_closest_preceding_index(idx, donts)

		if preceding_do >= preceding_dont {
			lhs, err_l := strconv.Atoi(string([]rune(input[m[2]:m[3]])))
			rhs, err_r := strconv.Atoi(string([]rune(input[m[4]:m[5]])))
			if err_l != nil || err_r != nil {
				panic("String conversion failed")
			}

			sum += lhs * rhs
		}
	}

	return sum
}

func find_closest_preceding_index(x int, indices []int) int {
	for i := range indices {
		diff := x - indices[i]

		if diff < 0 && i == 0 {
			return -1
		}

		if diff < 0 && i > 0 {
			return indices[i-1]
		}

		if diff > 0 && i == len(indices)-1 {
			return indices[i]
		}
	}

	return -1
}

func find_re_indices(input string, regex string) (indices []int) {
	re := regexp.MustCompile(regex)
	matches := re.FindAllStringIndex(input, -1)
	for _, match := range matches {
		indices = append(indices, match[0])
	}

	return
}
