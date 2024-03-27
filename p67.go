package main

import "fmt"

func main() {

	fmt.Println(winnable([]int{3, 3, 1, 0, 2, 0, 1}))
	fmt.Println(winnable([]int{3, 2, 0, 0, 2, 0, 1}))
	fmt.Println(winnable([]int{2, 4, 1, 1, 2, 0, 2, 3}))
	fmt.Println(canReachEnd([]int{3, 3, 1, 0, 2, 0, 1}))
	fmt.Println(canReachEnd([]int{3, 2, 0, 0, 2, 0, 1}))
	fmt.Println(canReachEnd([]int{2, 4, 1, 1, 2, 0, 2, 3}))

}

func winnable(in []int) bool {
	nextIndexes := []int{}
	nextIndexes = append(nextIndexes, 0)

	// bruteforce
	for len(nextIndexes) != 0 {
		i := nextIndexes[0]
		// can we reach the end from there ?
		if in[i]+i >= len(in)-1 {
			return true
		}
		// find squareToAdvance
		next := []int{}
		for j := i + 1; j-(i+1) < in[i] && j < len(in); j++ {
			if in[j] != 0 {
				next = append(next, j)
			}
		}
		nextIndexes = append(nextIndexes, next...)
		nextIndexes = nextIndexes[1:]
	}

	return false

}

// book solution, much more elegant
func canReachEnd(in []int) bool {
	furthest := 0
	for i := 0; i <= furthest && furthest < len(in); i++ {
		furthest = max(furthest, in[i]+i)
	}
	return furthest >= len(in)
}
