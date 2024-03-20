package main

import "fmt"

// 5.6 compute x/y using only addition, substraction, shifting
func main() {
	fmt.Println(bruteForce(15, 5))
}

func bruteForce(x, y int) int {
	result := 0

	for x >= y {
		result++
		x -= y
	}

	return result
}
