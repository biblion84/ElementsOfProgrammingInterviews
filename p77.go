package main

import (
	"fmt"
	"math/rand"
)

type pair struct {
	a, b int
}

func main() {

	total := map[pair]int{}

	for i := 0; i < 1_000_000; i++ {
		sample := sampling([]int{1, 2, 3, 4}, 2)
		total[pair{a: sample[0], b: sample[1]}]++
	}

	for p, t := range total {
		fmt.Printf("{%d,%d} : %d\n", p.a, p.b, t)
	}

}

func sampling(a []int, size int) []int {
	// NOTE: all permutation are not of equal probability, see 6.13 p80
	// 1 - scramble the array
	// 2 - return the subarray [:size]

	for i := 0; i < len(a); i++ {
		permutation := rand.Intn(len(a)-i) + i
		a[i], a[permutation] = a[permutation], a[i]
	}

	return a[:size]
}
