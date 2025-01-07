package day10

import (
	"strconv"
	"strings"

	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	emap, trailheads := readElevations(input)
	var numTrails int
	for _, th := range trailheads {
		summits := make(map[vec.Coord]bool)
		elev := emap[th]
		getSummits(th, elev, emap, &summits)
		numTrails += len(summits)
	}
	return numTrails
}

func Part2(input string) any {
	emap, trailheads := readElevations(input)
	var numTrails int
	for _, th := range trailheads {
		elev := emap[th]
		numTrails += getNumTrails(th, elev, emap)
	}

	return numTrails
}

type ElevMap map[vec.Coord]int

func getNumTrails(c vec.Coord, elev int, emap ElevMap) (numTrails int) {
	if elev == 9 {
		return 1
	}

	for _, n := range getNeighboringPaths(c, elev, emap) {
		nElev := emap[n]
		numTrails += getNumTrails(n, nElev, emap)
	}

	return numTrails
}

func getSummits(c vec.Coord, elev int, emap ElevMap, summits *map[vec.Coord]bool) {
	if elev == 9 {
		(*summits)[c] = true
	}

	for _, n := range getNeighboringPaths(c, elev, emap) {
		nElev := emap[n]
		getSummits(n, nElev, emap, summits)
	}

	return
}

func getNeighboringPaths(c vec.Coord, elev int, emap ElevMap) (neighbors []vec.Coord) {
	for _, d := range vec.AllDirections() {
		nCoord := c.Add(d)
		if nElev, ok := emap[nCoord]; ok && nElev == elev+1 {
			neighbors = append(neighbors, nCoord)
		}
	}

	return neighbors
}

func readElevations(input string) (ElevMap, []vec.Coord) {
	emap := make(ElevMap, len(input))
	var trailheads []vec.Coord
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			elev, err := strconv.Atoi(string(char))
			if err != nil {
				panic("Cannot convert to int")
			}

			coord := vec.Coord{X: i, Y: j}
			emap[coord] = elev
			if elev == 0 {
				trailheads = append(trailheads, coord)
			}
		}
	}

	return emap, trailheads
}
