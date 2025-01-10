package day24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) any {
	initValues, gates := readInput(input)

	var zs []string
	for k := range gates {
		if k[0] == 'z' {
			zs = append(zs, k)
		}
	}
	slices.Sort(zs)
	slices.Reverse(zs)

	var bits string
	for _, z := range zs {
		bits += fmt.Sprintf("%d", getValue(z, initValues, gates))
	}

	out, _ := strconv.ParseInt(bits, 2, 64)
	return out
}

func Part2(input string) any {
	fmt.Println("No solution yet for day 24, part 2")
	return 0
}

func getValue(wire string, initValues map[string]int, gates map[string]gate) int {
	if g, ok := gates[wire]; ok {
		return g.op(getValue(g.a, initValues, gates), getValue(g.b, initValues, gates))
	}

	if v, ok := initValues[wire]; ok {
		return v
	}

	fmt.Printf("Error: no match found for wire %s", wire)
	return -1
}

type gate struct {
	a, b string
	op   func(int, int) int
}

func bwand(a, b int) int {
	return a & b
}

func bwor(a, b int) int {
	return a | b
}

func bwxor(a, b int) int {
	return a ^ b
}

func readInput(input string) (initValues map[string]int, gates map[string]gate) {
	initValues = make(map[string]int)
	gates = make(map[string]gate)
	parts := strings.Split(input, "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		in := strings.Split(line, ": ")
		v, _ := strconv.Atoi(in[1])
		initValues[in[0]] = v
	}

	for _, line := range strings.Split(parts[1], "\n") {
		in := strings.Split(line, " ")
		var f func(int, int) int
		switch in[1] {
		case "AND":
			f = bwand
		case "OR":
			f = bwor
		case "XOR":
			f = bwxor
		default:
			fmt.Printf("Unknown operation %v", in[1])
		}
		gates[in[4]] = gate{a: in[0], b: in[2], op: f}
	}

	return initValues, gates
}
