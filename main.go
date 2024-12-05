package main

// Generate days.go, creating the list of days
//go:generate go run ./gen

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

type aocDay struct {
	Number int
	Part1  func(string) int
	Part2  func(string) int
}

// Only run from the root of the repo
func (day aocDay) GetInputPath(useExample bool) string {
	dayName := fmt.Sprintf("day%02d", day.Number)
	if useExample {
		return dayName + "/example.txt"
	} else {
		return dayName + "/input.txt"
	}
}

func main() {
	var dayNumber int
	flag.IntVar(&dayNumber, "d", 0, "The part of the puzzle")
	var part int
	flag.IntVar(&part, "p", 1, "The part of the puzzle")
	var useExample bool
	flag.BoolVar(&useExample, "e", false, "Use the example as input")
	flag.Parse()

	day, ok := days[dayNumber]
	if ok == false {
		fmt.Printf("Day %v does not exist", dayNumber)
		return
	}

	path := day.GetInputPath(useExample)
	inputBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}

	input := strings.TrimRight(string(inputBytes), "\n")
	var solution int
	if part == 1 {
		solution = day.Part1(input)
	} else {
		solution = day.Part2(input)
	}

	fmt.Println(solution)
}
