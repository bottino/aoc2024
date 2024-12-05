package day05

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

type RuleMap map[string](map[string]bool)

func Part1(input string) (solution int) {
	rules, updates := readInput(input)
	for _, update := range updates {
		if slices.IsSortedFunc(update, comparePages(rules)) {
			midValue, _ := strconv.Atoi(update[len(update)/2])
			solution += midValue
		}
	}

	return
}

func Part2(input string) (solution int) {
	rules, updates := readInput(input)
	for _, update := range updates {
		if slices.IsSortedFunc(update, comparePages(rules)) == false {
			slices.SortFunc(update, comparePages(rules))
			midValue, _ := strconv.Atoi(update[len(update)/2])
			solution += midValue
		}
	}

	return
}

func readInput(input string) (rules RuleMap, updates [][]string) {
	isFirstPart := true
	rules = make(RuleMap)
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

	return
}

func comparePages(rules RuleMap) func(string, string) int {
	return func(a, b string) int {
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
}
