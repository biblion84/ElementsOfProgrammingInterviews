package main

import "fmt"

func main() {

	for _, r := range permutation([]rune{'a', 'b', 'c', 'd'}, []int{2, 0, 1, 3}) {
		fmt.Print(string(r))
	}

}

func permutation(a []rune, perm []int) []rune {
	result := make([]rune, len(a))

	for i := 0; i < len(perm); i++ {
		result[perm[i]] = a[i]
	}

	return result
}

// The book give a solution in O(n) with O(1) space
// versus this O(n) with O(n) space
