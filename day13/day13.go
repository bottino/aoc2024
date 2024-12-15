package day13

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	eqs := readEqs(input)
	var coins int
	for _, eq := range eqs {
		na, nb := eq.Solve()
		if na < 0 || na > 100 || nb < 0 || nb > 100 {
			continue
		}

		coins += 3*na + nb
	}

	return coins
}

func Part2(input string) int {
	eqs := readEqs(input)
	var coins int
	for _, eq := range eqs {
		eq.xt += 1e13
		eq.yt += 1e13
		na, nb := eq.Solve()
		if na < 0 || nb < 0 {
			continue
		}

		coins += 3*na + nb
	}

	return coins
}

type Equation struct {
	xa, ya, xb, yb, xt, yt int
}

func (e *Equation) Solve() (na, nb int) {
	den := e.xa*e.yb - e.ya*e.xb
	// Ignore this case unless it becomes an issue
	if den == 0 {
		panic("div 0")
	}

	// get approximate solution in ints
	nb = (e.yt*e.xa - e.xt*e.ya) / den
	na = (e.xt - nb*e.xb) / e.xa

	// Check if we're perfectly on it
	if na*e.xa+nb*e.xb == e.xt && na*e.ya+nb*e.yb == e.yt {
		return na, nb
	} else {
		return 0, 0
	}
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
