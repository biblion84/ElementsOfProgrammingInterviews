package main

import "fmt"

func main() {

	fmt.Println(multiply([]int{2, 3}, []int{5, 4}))
	fmt.Println(multiply([]int{1, 9, 3, 7, 0, 7, 7, 2, 1}, []int{-7, 6, 1, 8, 3, 8, 2, 5, 7, 2, 8, 7}))
}

func multiply(a, b []int) []int {
	if len(a) > len(b) {
		a, b = b, a
	}
	negative := false
	if a[0] < 0 {
		negative = !negative
		a[0] *= -1
	}

	if b[0] < 0 {
		negative = !negative
		b[0] *= -1
	}

	maxSize := len(a) + len(b)
	result := make([]int, maxSize)
	resultIndex := 0

	for i := len(b) - 1; i >= 0; i-- {
		resultIndex = len(b) - i - 1
		carry := 0
		for j := len(a) - 1; j >= 0; j-- {
			r := (b[i] * a[j]) + carry
			result[resultIndex] += r
			carry = result[resultIndex] / 10
			result[resultIndex] = result[resultIndex] % 10
			resultIndex++
		}
		if carry != 0 {
			result[resultIndex] = carry
		}
	}

	if negative {
		result[resultIndex] *= -1
	}

	// reverse the array
	for i := 0; i < resultIndex/2; i++ {
		result[i], result[resultIndex-i] = result[resultIndex-i], result[i]
	}

	return result[:resultIndex+1]
}
