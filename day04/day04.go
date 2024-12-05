package day04

import (
	_ "embed"
	"fmt"
	"strings"
)

func Part1(input string) (xmas int) {
	mat := readInput(input)

	filters := []Mat{
		{{'X', 'M', 'A', 'S'}},
		{
			{'X', '.', '.', '.'},
			{'.', 'M', '.', '.'},
			{'.', '.', 'A', '.'},
			{'.', '.', '.', 'S'},
		},
	}

	allFilters := augmentFilters(filters)

	for _, filter := range allFilters {
		xmas += countFilterMatches(mat, filter)
	}

	return
}

func Part2(input string) (xmas int) {
	mat := readInput(input)

	filters := []Mat{
		{
			{'M', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'S'},
		},
	}

	allFilters := augmentFilters(filters)

	for _, filter := range allFilters {
		xmas += countFilterMatches(mat, filter)
	}

	return
}

func augmentFilters(filters []Mat) (allFilters []Mat) {
	for _, f := range filters {
		allFilters = append(allFilters, f, f.Rotate(90), f.Rotate(180), f.Rotate(-90))
	}

	return
}

func countFilterMatches(mat Mat, filter Mat) (matchCount int) {
	N, M := mat.Dims()
	n, m := filter.Dims()

	for i := range N - n + 1 {
		for j := range M - m + 1 {
			if isMatch(mat, filter, i, j) {
				matchCount++
			}
		}
	}

	return
}

func isMatch(mat Mat, filter Mat, x int, y int) bool {
	n, m := filter.Dims()
	for i := range n {
		for j := range m {
			if filter[i][j] != '.' && mat[x+i][y+j] != filter[i][j] {
				return false
			}
		}
	}

	return true
}

func readInput(input string) (mat Mat) {
	for _, line := range strings.Split(input, "\n") {
		var lines []rune
		for i := range line {
			lines = append(lines, rune(line[i]))
		}

		mat = append(mat, lines)
	}

	return mat
}

// Matrix operations
type Mat [][]rune

func (mat Mat) Rotate(angle int) (rotated Mat) {
	N, M := mat.Dims()

	if angle != 180 {
		rotated = newMat(M, N)
	} else {
		rotated = newMat(N, M)
	}

	for i := range N {
		for j := range M {
			switch angle {
			case 90:
				rotated[j][N-1-i] = mat[i][j]
			case 180:
				rotated[N-1-i][M-1-j] = mat[i][j]
			case -90:
				rotated[M-1-j][i] = mat[i][j]
			default:
				fmt.Printf("Unsupported angle %d", angle)
			}
		}
	}

	return
}

func (mat Mat) Dims() (n, m int) {
	return len(mat), len(mat[0])
}

func newMat(N int, M int) Mat {
	mat := make(Mat, N)
	for i := range mat {
		mat[i] = make([]rune, M)
	}

	return mat
}
