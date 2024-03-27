package main

import "fmt"

func main() {

	fmt.Println(removeDuplicates([]int{2, 3, 5, 5, 7, 11, 11, 13}))
	fmt.Println(removeDuplicates([]int{2, 2, 2, 2, 2, 3}))

}

func removeDuplicates(in []int) []int {

	validIndex := 0
	for i := 1; i < len(in); i++ {
		if in[i] != in[validIndex] {
			validIndex++
			in[validIndex] = in[i]
		}
	}
	return in[:validIndex+1]
}
