package main

import "fmt"

func main() {

	fmt.Println(increment([]int8{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}))

}

func increment(in []int8) []int8 {

	carry := int8(1)
	for i := len(in) - 1; i >= 0; i-- {
		in[i] += carry
		carry = 0
		if in[i] == 10 {
			in[i] = 0
			carry = 1
		}
	}
	// If we still have a carry at the end, we need to add a 1 in front
	if carry == 1 {
		in = append([]int8{1}, in...)
	}

	return in
}
