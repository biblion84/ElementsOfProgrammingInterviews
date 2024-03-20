package main

import "fmt"

// page 26
func main() {
	// test if x is a power of 2
	x := 0b00000100

	lowestBitSet := x & ^(x - 1)
	fmt.Println(x == lowestBitSet)
}

// right propagate the rightmost set bit in x
//func main() {
//
//	x := 0b00000100
//	lb := x & ^(x - 1)
//	fmt.Println((lb - 1) | x)
//
//}

// compute x modulo a power of two
//func main() {
//
//	x := 77
//	mod := 64
//
//	result := x & (mod - 1)
//	fmt.Println(result)
//
//}

// test if x is a power of 2
// func main() {
//	x := 0b00000100
//
//	lowestBitSet := x & ^(x - 1)
//	fmt.Println(x == lowestBitSet)
//}
