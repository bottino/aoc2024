package day10

import (
	"strconv"
	"strings"
)

func Part1(input string) any {
	emap, trailheads := readElevations(input)
	var numTrails int
	for _, th := range trailheads {
		summits := make(map[Coord]bool)
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

type Coord struct {
	X, Y int
}

type ElevMap map[Coord]int

func getNumTrails(c Coord, elev int, emap ElevMap) (numTrails int) {
	if elev == 9 {
		return 1
	}

	for _, n := range getNeighboringPaths(c, elev, emap) {
		nElev := emap[n]
		numTrails += getNumTrails(n, nElev, emap)
	}

	return numTrails
}

func getSummits(c Coord, elev int, emap ElevMap, summits *map[Coord]bool) {
	if elev == 9 {
		(*summits)[c] = true
	}

	for _, n := range getNeighboringPaths(c, elev, emap) {
		nElev := emap[n]
		getSummits(n, nElev, emap, summits)
	}

	return
}

func getNeighboringPaths(c Coord, elev int, emap ElevMap) (neighbors []Coord) {
	for _, d := range []Coord{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		nCoord := Coord{c.X + d.X, c.Y + d.Y}
		if nElev, ok := emap[nCoord]; ok && nElev == elev+1 {
			neighbors = append(neighbors, nCoord)
		}
	}

	return neighbors
}

func readElevations(input string) (ElevMap, []Coord) {
	emap := make(ElevMap, len(input))
	var trailheads []Coord
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			elev, err := strconv.Atoi(string(char))
			if err != nil {
				panic("Cannot convert to int")
			}

			coord := Coord{i, j}
			emap[coord] = elev
			if elev == 0 {
				trailheads = append(trailheads, coord)
			}
		}
	}

	return emap, trailheads
}
