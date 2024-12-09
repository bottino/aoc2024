package day08

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	nodeMap := readMap(input)
	antiNodeMap := make(map[Node]bool)
	for _, v := range nodeMap.NodeMap {
		pairs := getAllPairs(v)
		for _, p := range pairs {
			antiNode := getAntiNode(p[0], p[1])
			if isInBounds(antiNode, nodeMap) {
				antiNodeMap[antiNode] = true
			}
		}
	}

	return len(antiNodeMap)
}

func getAntiNode(a Node, b Node) Node {
	var u, v int = a.X - b.X, a.Y - b.Y
	return Node{a.X + u, a.Y + v, a.Freq}
}

func isInBounds(node Node, m Map) bool {
	return node.X >= 0 && node.X < m.N && node.Y > 0 && node.Y < m.M
}

func Part2(input string) (solution int) {
	fmt.Println("No solution yet for day 8, part 2")
	return
}

type Node struct {
	X, Y int
	Freq rune
}

type NodeMap map[rune][]Node

type Map struct {
	NodeMap
	N, M int
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

func readMap(input string) Map {
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

			n := Node{i, j, char}
			nodeMap[char] = append(v, n)
		}
	}

	return Map{nodeMap, len(lines), M}
}
