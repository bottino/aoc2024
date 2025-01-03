package day21

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bottino/aoc2024/graphs"
)

func Part1(input string) any {
	return getComplexity(input, 2)
}

func Part2(input string) any {
	return getComplexity(input, 25)
}

func getComplexity(input string, numRobots int) int {
	numPad := getShortestPaths(numKeys)
	arrowPad := getShortestPaths(arrowKeys)

	var complexity int
	memo := make(map[Mem]int)
	for _, code := range strings.Split(input, "\n") {
		seqs := getSeqs(code, numPad)
		minL := math.MaxInt
		for _, s := range seqs {
			shortest := shortestSeq(s, numRobots, arrowPad, &memo)
			if shortest < minL {
				minL = shortest
			}
		}

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

func getSeqs(code string, pad Pad) (seqs []string) {
	buildSeq(code, 0, "", pad, 'A', &seqs)
	return seqs
}

func buildSeq(code string, idx int, seq string, pad Pad, prev rune, result *[]string) {
	if idx == len(code) {
		(*result) = append(*result, seq)
		return
	}
	curr := rune(code[idx])
	paths := pad[Pair{prev, curr}]
	for _, p := range paths {
		buildSeq(code, idx+1, seq+p, pad, curr, result)
	}
}

func shortestSeq(seq string, depth int, pad Pad, memo *map[Mem]int) int {
	if depth == 0 {
		return len(seq)
	}

	if v, ok := (*memo)[Mem{seq, depth}]; ok {
		return v
	}

	var total int
	subSeqs := strings.SplitAfter(seq, "A")
	for _, sub := range subSeqs {
		seqs := getSeqs(sub, pad)
		minL := math.MaxInt
		for _, s := range seqs {
			length := shortestSeq(s, depth-1, pad, memo)
			if length < minL {
				minL = length
			}
		}

		total += minL
	}

	(*memo)[Mem{seq, depth}] = total
	return total
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

type Mem struct {
	seq   string
	depth int
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
