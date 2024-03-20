package main

import "fmt"

func main() {

	//fmt.Println(reverse(123456789))
	fmt.Println(reverse(-123))

}

func reverse(x int) int {
	result := 0
	for x != 0 {
		result *= 10
		result += x % 10
		x /= 10
	}

	return result
}
