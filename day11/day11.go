package day11

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input string) int {
	var numStones int
	stones := buildStoneMap(input)
	for i := 0; i < 25; i++ {
		stones, numStones = blink(stones)
	}

	return numStones
}

func Part2(input string) int {
	var numStones int
	stones := buildStoneMap(input)
	for i := 0; i < 75; i++ {
		stones, numStones = blink(stones)
	}

	return numStones
}

type Line map[int]int

func blink(stones Line) (newStones Line, numStones int) {
	newStones = make(Line, len(stones)*2)
	for k, v := range stones {

		if k == 0 {
			add(1, v, newStones)
			numStones += v
			continue
		}

		n := int(math.Log10(float64(k))) + 1
		if n%2 == 0 {
			a := int(float64(k) / math.Pow10(n/2))
			b := int(float64(k) - math.Pow10(n/2)*float64(a))
			add(a, v, newStones)
			add(b, v, newStones)
			numStones += 2 * v
			continue
		}

		add(2024*k, v, newStones)
		numStones += v
	}

	return newStones, numStones
}

func add(k int, v int, m Line) {
	if _, ok := m[k]; ok {
		m[k] += v
	} else {
		m[k] = v
	}
}

func buildStoneMap(input string) (stones Line) {
	stones = make(Line)
	for _, s := range strings.Split(input, " ") {
		d, _ := strconv.Atoi(s)
		add(d, 1, stones)
	}

	return stones
}
