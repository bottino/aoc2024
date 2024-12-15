package day12

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	garden := readGarden(input)
	N := len(garden)

	var regionId int
	garden[0][0].region = 0
	regions := []Region{{}}
	for i := 0; i < N; i++ {
		M := len(garden[i])
		for j := 0; j < M; j++ {
			lot := garden[i][j]
			var perimeter int

			for _, c := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
				nx, ny := i+c[0], j+c[1]
				// check bounds
				if nx < 0 || nx >= N || ny < 0 || ny >= M {
					perimeter++
					continue
				}

				n := garden[nx][ny]
				if n.plant == lot.plant {
					if n.region != -1 {
						lot.region = n.region
					}
				} else {
					perimeter++
				}
			}

			if lot.region == -1 {
				regionId++
				lot.region = regionId
				regions = append(regions, Region{})
			}

			fmt.Println(i, j, string(lot.plant), lot.region)

			regions[lot.region].area += 1
			regions[lot.region].perimeter += perimeter
		}
	}

	var solution int
	for _, r := range regions {
		fmt.Println(r)
		solution += r.area * r.perimeter
	}

	return solution
}

func Part2(input string) (solution int) {
	fmt.Println("No solution yet for day 12, part 2")
	return
}

type Garden [][]*Lot

type Lot struct {
	plant  rune
	region int
}

type Region struct {
	area      int
	perimeter int
}

type Coord struct {
	x, y int
}

func readGarden(input string) (garden Garden) {
	for _, line := range strings.Split(input, "\n") {
		row := make([]*Lot, len(line), len(line))
		for j, char := range line {
			row[j] = &Lot{char, -1}
		}
		garden = append(garden, row)
	}

	return garden
}
