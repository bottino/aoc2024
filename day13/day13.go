package day13

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	eqs := readEqs(input)
	var coins int
	for _, eq := range eqs {
		na, nb := eq.Solve()
		if na > 100 || nb > 100 {
			continue
		}

		coins += 3*na + nb
	}

	return coins
}

func Part2(input string) (solution int) {
	fmt.Println("No solution yet for day 13, part 2")
	return
}

type Equation struct {
	xa, ya, xb, yb, xt, yt int
}

func (e *Equation) Solve() (na, nb int) {
	den := e.xa*e.yb - e.ya*e.xb
	if den == 0 {
		if e.xt/e.xa == e.yt/e.ya {
			nb = e.xt / e.xa
			na = 0
		} else {
			na, nb = 0, 0
		}
		return na, nb
	}

	nb = (e.yt*e.xa - e.xt*e.ya) / den
	na = (e.xt - nb*e.xb) / e.xa
	return na, nb
}

func readEqs(input string) []Equation {
	var equations []Equation
	var eq Equation
	for _, line := range strings.Split(input, "\n") {
		re := regexp.MustCompile(`(.*): X(?:\+|=)(\d+), Y(?:\+|=)(\d+)`)
		matches := re.FindStringSubmatch(line)

		if len(matches) == 0 {
			continue
		}

		switch matches[1] {
		case "Button A":
			eq.xa, _ = strconv.Atoi(matches[2])
			eq.ya, _ = strconv.Atoi(matches[3])
		case "Button B":
			eq.xb, _ = strconv.Atoi(matches[2])
			eq.yb, _ = strconv.Atoi(matches[3])
		case "Prize":
			eq.xt, _ = strconv.Atoi(matches[2])
			eq.yt, _ = strconv.Atoi(matches[3])
			equations = append(equations, eq)
			eq = Equation{}
		}
	}

	return equations
}
