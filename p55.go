package main

import "fmt"

func main() {

	test := []int{0, 1, 7, 11, 121, 333, 2147447412, -1, 12, 100, 2147483647}

	for _, t := range test {
		fmt.Printf("%d: %t\n", t, isPalindrome(t))
	}

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverse(x)
}

// from p53
func reverse(x int) int {
	result := 0
	for x != 0 {
		result *= 10
		result += x % 10
		x /= 10
	}

	return result
}
