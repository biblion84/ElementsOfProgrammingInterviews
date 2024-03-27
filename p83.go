package main

import (
	"fmt"
	"math/rand"
)

func main() {

	result := nonUniformRandom([]int{3, 5, 7, 11}, []float32{9.0 / 18.0, 6.0 / 18.0, 2.0 / 18.0, 1.0 / 18.0}, 1_000_000)

	m := map[int]int{}

	for _, r := range result {
		m[r]++
	}

	for k, v := range m {
		fmt.Printf("%d : %d\n", k, v)
	}
}

// the sum of distribution must be 1
func nonUniformRandom(value []int, distribution []float32, nbToGenerate int) []int {

	result := make([]int, 0, nbToGenerate)

	for i := 0; i < nbToGenerate; i++ {
		r := rand.Float32()

		for j, d := range distribution {
			if r < d {
				result = append(result, value[j])
				break
			}
			r -= d
		}
	}

	return result
}
