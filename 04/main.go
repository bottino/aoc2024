package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func part1(input string) (words int) {
	var table [][]rune = read_input(input)

	var N int = len(table)
	var M int = len(table[0])

	// find line matches
	for _, line := range table {
		indices := find_matches(string(line))
		words += len(indices)
	}

	// find column matches
	columns := make([][]rune, M)
	for i := range columns {
		columns[i] = make([]rune, N)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			columns[i][j] = table[j][i]
		}
	}

	for _, col := range columns {
		indices := find_matches(string(col))
		words += len(indices)
	}

	// find diag matches
	for i := 0; i <= N-4; i++ {
		for j := 0; j <= M-4; j++ {
			square := get_square(table, i, j)
			words += get_square_diags(square)
		}
	}

	return
}

func get_square(table [][]rune, x int, y int) (square [4][4]rune) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			square[i][j] = table[x+i][y+j]
		}
	}

	return
}

func get_square_diags(square [4][4]rune) (words int) {
	var lr_diag string
	for i := 0; i < 4; i++ {
		lr_diag += string(square[i][i])
	}

	var rl_diag string
	for i := 0; i < 4; i++ {
		rl_diag += string(square[i][4-1-i])
	}

	if is_pattern(lr_diag) {
		words++
	}

	if is_pattern(rl_diag) {
		words++
	}

	return
}

func part2(input string) int {
	// part 2 here
	return 2
}

func find_matches(str string) (indices []int) {
	for i := 0; i <= len(str)-4; i++ { // Assuming ASCII chars
		if is_pattern(str[i : i+4]) {
			indices = append(indices, i)
		}
	}

	return indices
}

func is_pattern(chunk string) bool {
	return chunk == "XMAS" || chunk == "SAMX"
}

func read_input(input string) (table [][]rune) {
	for _, line := range strings.Split(input, "\n") {
		var lines []rune
		for i := range line {
			lines = append(lines, rune(line[i]))
		}

		table = append(table, lines)
	}

	return table
}
