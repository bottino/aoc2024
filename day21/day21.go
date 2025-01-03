package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bottino/aoc2024/graphs"
)

func Part1(input string) any {
	return 0
	// numPad := getShortestPaths(numKeys)
	// arrowPad := getShortestPaths(arrowKeys)
	//
	// var complexity int
	// for _, code := range strings.Split(input, "\n") {
	// 	seqs := process(code, []Pad{numPad, arrowPad, arrowPad})
	// 	minL := math.MaxInt
	// 	for _, s := range seqs {
	// 		if len(s) < minL {
	// 			minL = len(s)
	// 		}
	// 	}
	// 	numPart, err := strconv.Atoi(code[:3])
	// 	if err != nil {
	// 		fmt.Printf("Error when converting code %s", code)
	// 	}
	//
	// 	complexity += numPart * minL
	// }
	//
	// return complexity
}

func Part2(input string) any {
	numPad := getShortestPaths(numKeys)
	arrowPad := getShortestPaths(arrowKeys)

	pads := []Pad{numPad}
	for i := 0; i < 25; i++ {
		pads = append(pads, arrowPad)
	}

	var complexity int
	for _, code := range strings.Split(input, "\n") {
		seq := process(code, pads)
		minL := len(seq)
		numPart, err := strconv.Atoi(code[:3])
		if err != nil {
			fmt.Printf("Error when converting code %s", code)
		}

		complexity += numPart * minL
	}

	return complexity
}

func pathToString(path []rune, edges map[Pair]rune) string {
	out := ""
	for i := 0; i < len(path)-1; i++ {
		symbol, ok := edges[Pair{path[i], path[i+1]}]
		if !ok {
			panic("Edge not found")
		}

		out += string(symbol)
	}

	return out + "A"
}

func getShortestPaths(keys map[Coord]rune) Pad {
	numKp := graphs.New[rune]()
	edges := make(map[Pair]rune)
	for coord, key := range keys {
		for _, dir := range []Dir{left, right, up, down} {
			nCoord := coord.Add(dir.coord)
			if nKey, ok := keys[nCoord]; ok {
				numKp.AddEdge(key, nKey)
				edges[Pair{key, nKey}] = dir.symbol
			}
		}
	}

	shortestPaths := make(map[Pair][]string)
	for _, startNode := range numKp.Nodes() {
		_, prev := numKp.Dijkstra(startNode, graphs.UnitDist)
		for _, endNode := range numKp.Nodes() {
			pair := Pair{startNode, endNode}
			paths := numKp.GetAllShortestPaths(endNode, prev)
			for _, p := range paths {
				shortestPaths[pair] = append(shortestPaths[pair], pathToString(p, edges))
			}
		}
	}

	return shortestPaths
}

func process(code string, pads []Pad) (seq string) {
	memo := make(map[string]string, 1_000_000)
	seq = code
	for i, pad := range pads {
		fmt.Println(i)
		seq = processCode("A"+seq, pad, &memo)
	}

	return seq
}

func processCode(code string, pad Pad, memo *map[string]string) string {
	if s, ok := (*memo)[code]; ok {
		return s
	}

	if len(code) > 2 {
		left := code[:len(code)/2+1]
		right := code[len(code)/2:]
		leftSeq := processCode(left, pad, memo)
		rightSeq := processCode(right, pad, memo)
		(*memo)[left] = leftSeq
		(*memo)[right] = rightSeq
		return leftSeq + rightSeq
	}

	seqs, ok := pad[Pair{rune(code[0]), rune(code[1])}]
	if !ok {
		fmt.Printf("Error: couldn't read pair in pad {%s, %s}",
			string(code[0]), string(code[1]))
	}
	return seqs[0]
}

func concat(a []string, b []string) []string {
	out := make([]string, 0, len(a)*len(b))
	for j := 0; j < len(a); j++ {
		for k := 0; k < len(b); k++ {
			out = append(out, a[j]+b[k])
		}
	}

	return out
}

var (
	left  = Dir{'<', Coord{-1, 0}}
	right = Dir{'>', Coord{1, 0}}
	up    = Dir{'^', Coord{0, -1}}
	down  = Dir{'v', Coord{0, 1}}
)

func (lhs *Coord) Add(rhs Coord) Coord {
	return Coord{lhs.x + rhs.x, lhs.y + rhs.y}
}

type Pair struct {
	a, b rune
}

type Coord struct {
	x, y int
}

type Dir struct {
	symbol rune
	coord  Coord
}

type Pad map[Pair][]string

var numKeys = map[Coord]rune{
	{0, 0}: '7',
	{1, 0}: '8',
	{2, 0}: '9',
	{0, 1}: '4',
	{1, 1}: '5',
	{2, 1}: '6',
	{0, 2}: '1',
	{1, 2}: '2',
	{2, 2}: '3',
	{1, 3}: '0',
	{2, 3}: 'A',
}

var arrowKeys = map[Coord]rune{
	{1, 0}: '^',
	{2, 0}: 'A',
	{0, 1}: '<',
	{1, 1}: 'v',
	{2, 1}: '>',
}
