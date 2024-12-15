package day12

import (
	"fmt"
	"maps"
	"strings"
)

func Part1(input string) int {
	garden := readGarden(input)
	lots := maps.Clone(garden)
	regions := []Region{}
	var regionId int
	for len(lots) > 0 {
		lot := getFirstLot(lots)
		region := Region{id: regionId, coords: make(map[Coord]bool)}
		exploreRegion(lot, &region, garden)
		for k := range region.coords {
			delete(lots, k)
		}
		regionId++
		regions = append(regions, region)
	}

	var solution int
	for _, region := range regions {
		solution += region.area * region.perimeter
	}

	return solution
}

func Part2(input string) (solution int) {
	fmt.Println("No solution yet for day 12, part 2")
	return
}

func exploreRegion(lot Coord, region *Region, garden Garden) {
	region.area++
	region.coords[lot] = true
	for _, d := range []Coord{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		nCoord := Coord{lot.x + d.x, lot.y + d.y}
		nPlant, ok := garden[nCoord]

		// Handle off the map
		if !ok {
			region.perimeter++
			continue
		}

		// Handle already explored

		if garden[lot] == nPlant {
			if !region.coords[nCoord] {
				exploreRegion(nCoord, region, garden)
			}
		} else {
			region.perimeter++
		}
	}

	return
}

func getFirstLot(garden Garden) Coord {
	for k := range garden {
		return k
	}
	return Coord{0, 0}
}

type Coord struct {
	x, y int
}

type Region struct {
	id        int
	area      int
	perimeter int
	coords    map[Coord]bool
}

type Garden map[Coord]rune

func readGarden(input string) (garden Garden) {
	garden = make(Garden, len(input))
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			garden[Coord{i, j}] = char
		}
	}

	return garden
}
