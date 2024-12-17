package day09

import (
	"strconv"
)

func Part1(input string) any {
	var solution int
	memory, free := readInput(input)

	freeCount := 0
	freeId := 0
	for i := len(memory) - 1; i >= 0; i-- {
		v := memory[i]
		if v == -1 {
			freeCount++
		} else {
			memory[free[freeId]] = v
			memory[i] = -1
			freeCount++
			freeId++
		}

		if freeCount >= len(free) {
			break
		}
	}

	for i, v := range memory {
		if v == -1 {
			break
		}

		solution += i * v
	}

	return solution
}

func Part2(input string) any {
	chunks, free := readChunks(input)
	memory := chunksToInts(chunks)
	for i := len(chunks) - 1; i >= 0; i-- {
		c := chunks[i]
		if c.Id == -1 {
			continue
		}

		for _, f := range free {
			if c.Index > f.Index && c.Len <= f.Len {
				for j := 0; j < c.Len; j++ {
					memory[f.Index+j] = c.Id
					memory[c.Index+j] = -1
				}

				// Fill the free chunk from the right
				f.Index += c.Len
				f.Len -= c.Len
				break
			}
		}
	}

	var solution int
	for i, v := range memory {
		if v == -1 {
			continue
		}

		solution += i * v
	}

	return solution
}

type Chunk struct {
	Index int
	Len   int
	Id    int
}

func readChunks(input string) (memory []Chunk, free []*Chunk) {
	memory = make([]Chunk, 0, len(input))
	free = make([]*Chunk, 0, len(input))
	var index int
	for i, char := range input {
		blockL, _ := strconv.Atoi(string(char))
		chunk := Chunk{index, blockL, i / 2}

		if i%2 != 0 {
			chunk.Id = -1
			free = append(free, &chunk)
		}

		memory = append(memory, chunk)
		index += blockL
	}

	return memory, free
}

func chunksToInts(chunks []Chunk) []int {
	output := make([]int, 0, len(chunks)*10)
	for _, c := range chunks {
		for i := 0; i < c.Len; i++ {
			output = append(output, c.Id)
		}
	}
	return output
}

func readInput(input string) ([]int, []int) {
	output := make([]int, 0, len(input)*10)
	freeSpaces := make([]int, 0, len(input)*10)
	var memAddress int
	for i, char := range input {
		blockL, _ := strconv.Atoi(string(char))
		for j := 0; j < blockL; j++ {
			if i%2 == 0 {
				output = append(output, i/2)
			} else {
				output = append(output, -1)
				freeSpaces = append(freeSpaces, memAddress)
			}
			memAddress++
		}
	}

	return output, freeSpaces
}
