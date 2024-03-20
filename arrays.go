package main

import (
	"fmt"
	"slices"
)

// page 31
func main() {

	coins := []int{1, 1, 1, 5, 10, 25, 1, 1}

	slices.Sort(coins)

	maxValue := 0
	for _, c := range coins {
		if c > maxValue+1 {
			break
		}
		maxValue += c
	}

	fmt.Println(maxValue + 1)
}
