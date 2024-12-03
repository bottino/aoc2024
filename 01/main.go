package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
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
	var left, right []int = return_lists(input)
	slices.Sort(left)
	slices.Sort(right)

	var total_distance int
	for i := range left {
		total_distance += absInt(left[i] - right[i])
	}

	return total_distance
}

func part2(input string) int {
	var left, right []int = return_lists(input)
	var similarity int
	for i := range left {
		for j := range right {
			if left[i] == right[j] {
				similarity += left[i]
			}
		}
	}

	return similarity
}

func return_lists(input string) (left, right []int) {
	for _, line := range strings.Split(input, "\n") {
		re := regexp.MustCompile(`(\d+)\s+(\d+)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			left_value, err_l := strconv.Atoi(matches[1])
			right_value, err_r := strconv.Atoi(matches[2])
			if err_l != nil || err_r != nil {
				panic("String conversion failed")
			}

			left = append(left, left_value)
			right = append(right, right_value)
		} else {
			panic("No match found")
		}
	}

	return left, right
}

func absInt(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
