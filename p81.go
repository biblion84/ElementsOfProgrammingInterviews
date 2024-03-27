package main

import (
	"fmt"
	"math/rand"
)

type Permutation struct {
	a, b int
}

func main() {

	a := map[Permutation]int{}

	for i := 0; i < 100_000; i++ {
		p := randomSubset([]int{0, 1, 2}, 2)
		permutation := Permutation{p[0], p[1]}
		a[permutation]++
	}

	for k, v := range a {
		fmt.Printf("[%d %d] : %d\n", k.a, k.b, v)
	}
}

func randomSubset(in []int, size int) []int {
	return randomPermutation(in)[:size]
}

func randomPermutation(in []int) []int {
	for i := len(in) - 1; i >= 0; i-- {
		sample := rand.Intn(i + 1)
		in[i], in[sample] = in[sample], in[i]
	}
	return in
}
