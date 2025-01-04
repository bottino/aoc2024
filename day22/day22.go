package day22

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) any {
	var totSecrets int
	for _, secStr := range strings.Split(input, "\n") {
		secret, _ := strconv.Atoi(secStr)
		for i := 0; i < 2000; i++ {
			secret = getSecret(secret)
		}

		totSecrets += secret
	}

	return totSecrets
}

func Part2(input string) any {
	fmt.Println("No solution yet for day 22, part 2")
	return 0
}

type seq [4]int

func getSequences(init int, n int) map[seq]int {
	bananas, changes := getDiff(init, n)
	sequences := make(map[seq]int, len(changes)-4)
	// start from the end, because only the FIRST sequence counts
	for i := len(changes) - 1; i >= 3; i-- {
		s := seq{}
		for j := 0; j < 4; j++ {
			s[j] = changes[i-3+j]
		}
		sequences[s] = bananas[i+1]
	}

	return sequences
}

func getDiff(init int, n int) (bananas []int, changes []int) {
	secret := init
	first := init % 10
	bananas = append(bananas, first)
	prev := first
	for i := 1; i < n; i++ {
		secret = getSecret(secret)
		b := secret % 10
		bananas = append(bananas, b)
		changes = append(changes, b-prev)
		prev = b
	}
	return bananas, changes
}

func getSecret(secret int) int {
	x := secret * 64
	secret = mix(secret, x)
	secret = prune(secret)

	x = secret / 32
	secret = mix(secret, x)
	secret = prune(secret)

	x = secret * 2048
	secret = mix(secret, x)
	secret = prune(secret)
	return secret
}

func mix(secret int, value int) int {
	return value ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}
