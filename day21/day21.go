package day21

import (
	"fmt"

	"github.com/bottino/aoc2024/graphs"
)

func Part1(input string) any {
	// codes := strings.Split(input, "\n")

	// numPad := getShortestPaths(numKeys)

	// for _, code := range codes {
	//
	// }

	return 0
}

func Part2(input string) any {
	fmt.Println("No solution yet for day 21, part 2")
	return 0
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
			if startNode == endNode {
				continue
			}

			pair := Pair{startNode, endNode}
			paths := numKp.GetAllShortestPaths(endNode, prev)
			for _, p := range paths {
				shortestPaths[pair] = append(shortestPaths[pair], pathToString(p, edges))
			}
		}
	}

	return shortestPaths
}

func processCode(code string, pad Pad) (seqs []string) {
	code = "A" + code
	for i := 0; i < len(code)-1; i++ {
		paths, ok := pad[Pair{rune(code[i]), rune(code[i+1])}]
		if !ok {
			panic("shouldn't error here")
		}

		newSeqs := make([]string, 0, len(paths)*len(seqs))
		fmt.Println(seqs, paths)

		if len(seqs) == 0 {
			seqs = append(seqs, paths...)
			continue
		}

		for j := 0; j < len(seqs); j++ {
			for k := 0; k < len(paths); k++ {
				newSeqs = append(newSeqs, seqs[j]+paths[k])
			}
		}

		seqs = newSeqs
	}

	return seqs
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
