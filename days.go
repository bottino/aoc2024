package main

import (
	"github.com/bottino/aoc2024/day01"
	"github.com/bottino/aoc2024/day02"
	"github.com/bottino/aoc2024/day03"
	"github.com/bottino/aoc2024/day04"
	"github.com/bottino/aoc2024/day05"
)

var days = map[int]aocDay{
	1: {1, day01.Part1, day01.Part2},
	2: {2, day02.Part1, day02.Part2},
	3: {3, day03.Part1, day03.Part2},
	4: {4, day04.Part1, day04.Part2},
	5: {5, day05.Part1, day05.Part2},
}
