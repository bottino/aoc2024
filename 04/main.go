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

const L1 = 4

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
	var table [][]rune = readInput(input)

	var N int = len(table)
	var M int = len(table[0])

	// find line matches
	for _, line := range table {
		indices := findMatches(string(line))
		words += len(indices)
	}

	// find column matches
	columns := make2dMat(M, N)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			columns[i][j] = table[j][i]
		}
	}

	for _, col := range columns {
		indices := findMatches(string(col))
		words += len(indices)
	}

	// find diag matches
	for i := 0; i <= N-L1; i++ {
		for j := 0; j <= M-L1; j++ {
			square := getSquare(table, i, j, L1)
			words += getDiagWords(square)
		}
	}

	return
}

func part2(input string) int {
	// part 2 here
	return 2
}

func make2dMat(N int, M int) [][]rune {
	mat := make([][]rune, N)
	for i := range mat {
		mat[i] = make([]rune, M)
	}

	return mat
}

func getSquare(table [][]rune, x int, y int, L int) (square [][]rune) {
	square = make2dMat(L, L)
	for i := 0; i < L; i++ {
		for j := 0; j < L; j++ {
			square[i][j] = table[x+i][y+j]
		}
	}

	return
}

func getSquareDiags(square [][]rune) (lrDiag string, rlDiag string) {
	for i := 0; i < L1; i++ {
		lrDiag += string(square[i][i])
	}

	for i := 0; i < L1; i++ {
		rlDiag += string(square[i][L1-1-i])
	}

	return lrDiag, rlDiag
}

func getDiagWords(square [][]rune) (words int) {
	lrDiag, rlDiag := getSquareDiags(square)

	if isXmasPattern(lrDiag) {
		words++
	}

	if isXmasPattern(rlDiag) {
		words++
	}

	return
}

func findMatches(str string) (indices []int) {
	for i := 0; i <= len(str)-L1; i++ { // Assuming ASCII chars
		if isXmasPattern(str[i : i+L1]) {
			indices = append(indices, i)
		}
	}

	return indices
}

func isXmasPattern(chunk string) bool {
	return chunk == "XMAS" || chunk == "SAMX"
}

func readInput(input string) (table [][]rune) {
	for _, line := range strings.Split(input, "\n") {
		var lines []rune
		for i := range line {
			lines = append(lines, rune(line[i]))
		}

		table = append(table, lines)
	}

	return table
}
