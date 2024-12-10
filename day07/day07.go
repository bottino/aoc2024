package day07

import (
	"math"
	"strconv"
	"strings"
)

var add Operator = func(a int, b int) int {
	return a + b
}
var mul Operator = func(a int, b int) int {
	return a * b
}
var concat Operator = func(a int, b int) int {
	numDigits := int(math.Log10(float64(b))) + 1
	return a*int(math.Pow(10, float64(numDigits))) + b
}

func Part1(input string) int {
	operators := []Operator{add, mul}
	return solve(input, operators)
}

func Part2(input string) int {
	operators := []Operator{add, mul, concat}
	return solve(input, operators)
}

func solve(input string, operators []Operator) (solution int) {
	eqs := readEqs(input)

	for _, eq := range eqs {
		solved := recurseSolve(eq.Res, eq.Operands[0], eq.Operands[1:], operators)
		if solved {
			solution += eq.Res
		}
	}

	return solution
}

func recurseSolve(res int, curr int, operands []int, operators []Operator) bool {
	if curr > res {
		return false
	}
	if len(operands) == 0 {
		return res == curr
	}

	for _, op := range operators {
		solved := recurseSolve(res, op(curr, operands[0]), operands[1:], operators)
		if solved {
			return true
		}
	}

	return false
}

type Equation struct {
	Res      int
	Operands []int
}

type Operator func(int, int) int

func readEqs(input string) (equations []Equation) {
	for _, line := range strings.Split(input, "\n") {
		s := strings.Split(line, ": ")
		result, _ := strconv.Atoi(s[0])
		var ops []int
		for _, s := range strings.Split(s[1], " ") {
			o, _ := strconv.Atoi(s)
			ops = append(ops, o)
		}
		equations = append(equations, Equation{result, ops})
	}

	return equations
}
