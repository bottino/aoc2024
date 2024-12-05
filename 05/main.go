package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type RuleMap map[string](map[string]bool)

var rules RuleMap

func part1(input string) int {
	isFirstPart := true
	rules = make(RuleMap)
	var updates [][]string
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			isFirstPart = false
			continue
		}

		if isFirstPart {
			r := strings.Split(line, "|")
			if _, ok := rules[r[0]]; ok == false {
				rules[r[0]] = make(map[string]bool)
			}
			rules[r[0]][r[1]] = true
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	var sumMidValues int
	for _, update := range updates {
		if slices.IsSortedFunc(update, comparePages) {
			midValue, _ := strconv.Atoi(update[len(update)/2])
			sumMidValues += midValue
		}
	}

	return sumMidValues
}

func part2(input string) int {
	// part 2 here
	return 2
}

func comparePages(a string, b string) int {
	// b must be after a
	if _, ok := rules[a][b]; ok {
		return -1
	}

	// b must be before a
	if _, ok := rules[b][a]; ok {
		return 1
	}

	// incomparable
	return 0
}
