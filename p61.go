package main

import "fmt"

func main() {

	fmt.Println(dutchFlag([]int{0, 1, 2, 0, 2, 1, 1}, 3))
	fmt.Println(dutchFlag([]int{0, 1, 2, 0, 2, 1, 1}, 2))
}

func dutchFlag(a []int, pivotIndex int) []int {
	pivot := a[pivotIndex]

	smaller := 0
	for i := 0; i < len(a); i++ {
		if a[i] < pivot {
			a[smaller], a[i] = a[i], a[smaller]
			smaller++
		}
	}

	bigger := len(a) - 1
	for i := len(a) - 1; i >= smaller; i-- {
		if a[i] > pivot {
			a[bigger], a[i] = a[i], a[bigger]
			bigger--
		}
	}

	return a
}
