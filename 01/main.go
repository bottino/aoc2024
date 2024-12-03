package main

import (
	_ "embed"
	"fmt"
	"log"
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
	example = strings.TrimRight(input, "\n")
	input = strings.TrimRight(input, "\n")
}

func main() {
	var left []int
	var right []int
	for _, line := range strings.Split(input, "\n") {
		re := regexp.MustCompile(`(\d+)\s+(\d+)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			left_value, _ := strconv.Atoi(matches[1])
			right_value, _ := strconv.Atoi(matches[2])
			left = append(left, left_value)
			right = append(right, right_value)
		} else {
			log.Fatal("No match found")
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	var total_distance int
	for i := range left {
		total_distance += absInt(left[i] - right[i])
	}
	fmt.Println(total_distance)
}

func absInt(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
