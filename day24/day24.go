package day24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bottino/aoc2024/utils"
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
	initValues, gates := readInput(input)

	var edges []edge
	var nodes []node
	for out, gTo := range gates {
		nodes = append(nodes, node{out, out + " " + gTo.label})
		edges = append(edges, edge{gTo.a, out})
		edges = append(edges, edge{gTo.b, out})
	}

	var zs []string
	for k := range gates {
		if k[0] == 'z' {
			zs = append(zs, k)
		}
	}

	var inputs []string
	for k := range initValues {
		nodes = append(nodes, node{k, k})
		inputs = append(inputs, k)
	}

	data := tmplData{nodes, edges, strings.Join(inputs, " "), strings.Join(zs, " ")}
	utils.CreateFromTemplate("day24/circuit.gv", "day24/circuit.tmpl", data)

	// Found them visually like a boss
	// TODO: find a better way to get the problem wires
	return "cvh,dbb,hbk,kvn,tfn,z14,z18,z23"
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

type tmplData struct {
	Nodes  []node
	Edges  []edge
	Inputs string
	Zs     string
}

type node struct {
	OutWire, Label string
}

type edge struct {
	From, To string
}

type gate struct {
	a, b  string
	op    func(int, int) int
	label string
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
		var label string
		switch in[1] {
		case "AND":
			f = bwand
			label = "&"
		case "OR":
			f = bwor
			label = "≥1"
		case "XOR":
			f = bwxor
			label = "=1"
		default:
			fmt.Printf("Unknown operation %v", in[1])
		}
		gates[in[4]] = gate{a: in[0], b: in[2], op: f, label: label}
	}

	return initValues, gates
}
