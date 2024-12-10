package day07

import (
	"math"
	"strconv"
	"strings"
	"sync"
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

	c := make(chan int, len(eqs))
	wg := sync.WaitGroup{}
	for _, eq := range eqs {
		wg.Add(1)
		go SolveEq(eq, operators, c, &wg)
	}

	wg.Wait()
	close(c)

	for res := range c {
		solution += res
	}

	return solution
}

func SolveEq(eq Equation, operators []Operator, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	combinations := genCombinations(operators, len(eq.Operands)-1)
	for _, ops := range combinations {
		res := eq.Operands[0]
		for i := 0; i < len(eq.Operands)-1; i++ {
			res = ops[i](res, eq.Operands[i+1])
			if res > eq.Res {
				break
			}
		}

		if res == eq.Res {
			c <- res
			return
		}
	}
}

type Equation struct {
	Res      int
	Operands []int
}

type Operator func(int, int) int

func genCombinations[T any](items []T, n int) (combinations [][]T) {
	var helper func(curr []T, length int)
	helper = func(curr []T, length int) {
		if length == 0 {
			c := make([]T, len(curr))
			copy(c, curr)
			combinations = append(combinations, c)
			return
		}

		for _, item := range items {
			helper(append(curr, item), length-1)
		}
	}

	helper([]T{}, n)
	return combinations
}

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
