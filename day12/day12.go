package day12

import (
	"maps"
	"strings"

	"github.com/bottino/aoc2024/vec"
)

func Part1(input string) any {
	regions := processRegions(input)
	var solution int
	for _, region := range regions {
		solution += region.area * region.perimeter
	}

	return solution
}

func Part2(input string) any {
	regions := processRegions(input)
	var solution int
	for _, region := range regions {
		solution += region.area * region.numSides
	}

	return solution
}

func processRegions(input string) []Region {
	garden := readGarden(input)
	lots := maps.Clone(garden)
	regions := []Region{}
	var regionId int
	for len(lots) > 0 {
		lot := getFirstLot(lots)
		region := Region{
			id:    regionId,
			lots:  make(map[vec.Coord]bool),
			sides: make(map[Side]bool),
		}
		exploreRegion(lot, &region, garden)
		for k := range region.lots {
			delete(lots, k)
		}

		var numSides int
		for side := range region.sides {
			adj := FindAdjSide(side)
			if !region.sides[adj] {
				numSides++
			}
		}

		region.numSides = numSides
		regionId++
		regions = append(regions, region)
	}

	return regions
}

func FindAdjSide(side Side) Side {
	dir := side.dir
	lot := side.lot
	switch dir {
	case vec.Up:
		return Side{vec.Up, lot.Add(vec.Right)}
	case vec.Right:
		return Side{vec.Right, lot.Add(vec.Down)}
	case vec.Down:
		return Side{vec.Down, lot.Add(vec.Left)}
	case vec.Left:
		return Side{vec.Left, lot.Add(vec.Up)}
	default:
		panic("went wrong")
	}
}

func exploreRegion(lot vec.Coord, region *Region, garden Garden) {
	region.area++
	region.lots[lot] = true
	for _, d := range vec.AllDirections() {
		nCoord := lot.Add(d)
		nPlant, ok := garden[nCoord]
		side := Side{d, lot}

		// Handle off the map
		if !ok {
			region.perimeter++
			region.sides[side] = true
			continue
		}

		if garden[lot] == nPlant {
			if !region.lots[nCoord] {
				exploreRegion(nCoord, region, garden)
			}
		} else {
			region.perimeter++
			region.sides[side] = true
		}
	}

	return
}

func getFirstLot(garden Garden) vec.Coord {
	for k := range garden {
		return k
	}
	return vec.Coord{X: 0, Y: 0}
}

type Side struct {
	dir vec.Coord
	lot vec.Coord
}

type Region struct {
	id        int
	area      int
	perimeter int
	numSides  int
	lots      map[vec.Coord]bool
	sides     map[Side]bool
}

type Garden map[vec.Coord]rune

func readGarden(input string) (garden Garden) {
	garden = make(Garden, len(input))
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			garden[vec.Coord{X: i, Y: j}] = char
		}
	}

	return garden
}
