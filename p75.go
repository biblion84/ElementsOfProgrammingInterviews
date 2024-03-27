package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println(nextPermutation([]int{1, 0, 3, 2}))
	fmt.Println(nextPermutation([]int{1, 5, 3, 4}))
	fmt.Println(nextPermutation([]int{1, 3, 4, 5}))

}

func nextPermutation(in []int) []int {
	k := len(in) - 2
	for ; k >= 0 && in[k] >= in[k+1]; k-- {
	}
	if k == -1 {

	}

	for i := len(in) - 1; i >= 0; i-- {
		if in[i] > in[k] {
			in[i], in[k] = in[k], in[i]
			break
		}
	}

	slices.Reverse(in[k+1:])

	return in
}
