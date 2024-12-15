package day12

import (
	"maps"
	"strings"
)

func Part1(input string) int {
	regions := processRegions(input)
	var solution int
	for _, region := range regions {
		solution += region.area * region.perimeter
	}

	return solution
}

func Part2(input string) int {
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
			lots:  make(map[Coord]bool),
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

var (
	up    = Coord{-1, 0}
	down  = Coord{1, 0}
	left  = Coord{0, -1}
	right = Coord{0, 1}
)

func FindAdjSide(side Side) Side {
	dir := side.dir
	lot := side.lot
	switch dir {
	case up:
		return Side{up, lot.Add(right)}
	case right:
		return Side{right, lot.Add(down)}
	case down:
		return Side{down, lot.Add(left)}
	case left:
		return Side{left, lot.Add(up)}
	default:
		panic("went wrong")
	}
}

func (lhs *Coord) Add(rhs Coord) Coord {
	return Coord{lhs.x + rhs.x, lhs.y + rhs.y}
}

func addSides(side Side, region *Region) {
}

func exploreRegion(lot Coord, region *Region, garden Garden) {
	region.area++
	region.lots[lot] = true
	for _, d := range []Coord{up, down, left, right} {
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

func getFirstLot(garden Garden) Coord {
	for k := range garden {
		return k
	}
	return Coord{0, 0}
}

type Coord struct {
	x, y int
}

type Side struct {
	dir Coord
	lot Coord
}

type Region struct {
	id        int
	area      int
	perimeter int
	numSides  int
	lots      map[Coord]bool
	sides     map[Side]bool
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
