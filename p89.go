package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(rotate(grid))
}

func rotate(in [][]int) [][]int {
	result := make([][]int, len(in))
	for i := range result {
		result[i] = make([]int, 0, len(in[i]))
	}

	for x := 0; x < len(in); x++ {
		for y := len(in) - 1; y >= 0; y-- {
			result[x] = append(result[x], in[y][x])
		}
	}

	return result
}
