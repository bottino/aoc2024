package day08

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	return Solve(input, true)
}

func Part2(input string) int {
	return Solve(input, false)
}

func Solve(input string, part1 bool) int {
	freqMap := readMap(input)
	freqMap.AnodeMap = make(map[Node]bool)

	fmt.Println(freqMap)

	for _, v := range freqMap.NodeMap {
		pairs := getAllPairs(v)
		for _, p := range pairs {
			var antiNodes []Node
			if part1 {
				antiNodes = getAntiNodesP1(p[0], p[1], freqMap)
			} else {
				antiNodes = getAntiNodesP2(p[0], p[1], freqMap)
			}

			for _, n := range antiNodes {
				freqMap.AnodeMap[n] = true
			}
		}
	}

	fmt.Println(freqMap)

	return len(freqMap.AnodeMap)
}

func getAntiNodesP2(a Node, b Node, fMap FreqMap) (anodes []Node) {
	var u, v int = b.X - a.X, b.Y - a.Y
	f := 0
	n := b
	for fMap.isInBounds(n) {
		anodes = append(anodes, n)
		f++
		n = Node{b.X + u*f, b.Y + v*f}
	}

	return anodes
}

func getAntiNodesP1(a Node, b Node, fMap FreqMap) []Node {
	var u, v int = b.X - a.X, b.Y - a.Y
	n := Node{b.X + u, b.Y + v}
	if fMap.isInBounds(n) {
		return []Node{n}
	} else {
		return []Node{}
	}
}

func (m *FreqMap) isInBounds(node Node) bool {
	return node.X >= 0 && node.X < m.N && node.Y >= 0 && node.Y < m.M
}

type Node struct {
	X, Y int
}

type NodeMap map[rune][]Node

type FreqMap struct {
	NodeMap
	AnodeMap map[Node]bool
	N, M     int
}

func (m FreqMap) String() string {
	fullMap := make([][]rune, m.N)
	for i := range m.N {
		fullMap[i] = make([]rune, m.M)
		for j := range m.M {
			fullMap[i][j] = '.'
		}
	}

	for freq, v := range m.NodeMap {
		for _, n := range v {
			fullMap[n.X][n.Y] = freq
		}
	}

	for k := range m.AnodeMap {
		fullMap[k.X][k.Y] = '#'
	}

	var output string
	for i := range m.N {
		output += string(fullMap[i])
		output += "\n"
	}
	return output
}

func getAllPairs(nodes []Node) (allPairs [][2]Node) {
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			if i == j {
				continue
			}

			allPairs = append(allPairs, [2]Node{nodes[i], nodes[j]})
		}
	}

	return allPairs
}

func readMap(input string) FreqMap {
	nodeMap := make(NodeMap)
	var M int
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		M = len(line)
		for j, char := range line {
			if char == '.' {
				continue
			}

			v, ok := nodeMap[char]
			if ok == false {
				nodeMap[char] = []Node{}
			}

			n := Node{i, j}
			nodeMap[char] = append(v, n)
		}
	}

	return FreqMap{nodeMap, map[Node]bool{}, len(lines), M}
}
