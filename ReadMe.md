# Advent of Code 2024, 100% Go

This repository contains my solutions to the
[Advent of Code 2024 challenge](https://adventofcode.com/2024). This year, I
decided to learn [Go](https://go.dev/).

Because of a busy holiday schedule, I only managed to finish the challenge on
January 10, 2024.

## Solutions

All solutions are entirely my own. In a few cases in the early days, I first got
to a solution on my own, and then implemented another solution based on a
friend's suggestion.

For algorithms, I used my faithful copy of
["Algorithms" by Sedgewick and Wayne](https://algs4.cs.princeton.edu/home/) as a
reference. I also had to Google a specific algorithm to solve day 23.

Since the goal was to learn Go (and some useful algorithms along the way), I
didn't use any genAI tools. I did use a language server however, and had to look
up a few specifics in the Go documentation, for instance how the `heap` package
worked for my Disjkstra implementation

All the solutions use only the standard library. The only dependency is
[go-cmp](https://github.com/google/go-cmp), that I used for unit tests. I wrote
a small helper package for graph operations, and another for working with 2D
grids.

Some solutions are not optimal. For instance, my solution for day 18, part 2
takes several seconds. Also, I had to resort to visual inspection to solve
day 23.

## Structure

Individual inputs to Advent of Code problems shouldn't be distributed, so they
live in the "inputs" submodule that links to a private repository.

Each day is a separate package with two exposed functions: `Part1` and `Part2`,
with signature `func (string) any`. It would have been slightly more optimal and
idiomatic to have `func (io.Reader) any` functions, but I used `string` input
for simplicity.

`main.go` can run the solution for any day, or for all days when using the
`--all` flag. To import all packages in the `main` package, we use code
generation (see the "gen" package).

The "start" package builds a command to create a new AOC day.
