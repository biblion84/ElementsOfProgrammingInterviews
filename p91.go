package main

import "fmt"

func main() {

	fmt.Println(pascalTriangle(10))

}

func pascalTriangle(x int) [][]int {
	result := make([][]int, 0, x)
	result = append(result, []int{1})

	for i := 1; i < x; i++ {
		nextRow := make([]int, len(result[i-1])+1)
		for j := 0; j < len(nextRow); j++ {
			if j == 0 || j == len(nextRow)-1 {
				nextRow[j] = 1
			} else {
				nextRow[j] = result[i-1][j] + result[i-1][j-1]
			}
		}

		result = append(result, nextRow)
	}

	return result
}
