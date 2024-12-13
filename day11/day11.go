package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(input string) int {
	stones := buildStoneMap(input)
	for i := 0; i < 25; i++ {
		newStones := blink(stones)
		stones = append(stones, newStones...)
	}

	return len(stones)
}

func Part2(input string) (solution int) {
	stones := buildStoneMap(input)
	for i := 0; i < 75; i++ {
		newStones := blink(stones)
		stones = append(stones, newStones...)
		fmt.Println(i+1, len(stones))
	}

	return len(stones)
}

type Line map[int]int

func blink(stones []int) (newStones []int) {
	for i, v := range stones {

		if v == 0 {
			stones[i] = 1
			continue
		}

		n := numDigits(v)
		if n%2 == 0 {
			a := int(float64(v) / math.Pow10(n/2))
			b := int(float64(v) - math.Pow10(n/2)*float64(a))
			stones[i] = a
			newStones = append(newStones, b)
			continue
		}

		stones[i] = 2024 * v
	}

	return newStones
}

func numDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func buildStoneMap(input string) (stones []int) {
	for _, s := range strings.Split(input, " ") {
		d, _ := strconv.Atoi(s)
		stones = append(stones, d)
	}

	return stones
}
