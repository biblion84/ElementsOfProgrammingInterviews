package main

import "fmt"

var PARITY_LOOKUP_TABLE [256]int

// 5.1
func main() {
	fmt.Println(parity(0b1011))
	fmt.Println(parity(0b10001000))
	for i := range PARITY_LOOKUP_TABLE {
		PARITY_LOOKUP_TABLE[i] = parity(int64(i))
	}
	fmt.Println(parityLookup(0b1011))
	fmt.Println(parityLookup(0b10001000))
	fmt.Println(parityXor(0b1011))
	fmt.Println(parityXor(0b10001000))
}

// Simple brute force way
func parity(x int64) int {
	popCount := 0
	for x != 0 {
		if x&1 == 1 {
			popCount++
		}
		x >>= 1
	}
	return popCount % 2
}

// using lookup table
func parityLookup(x int64) int {
	parity := 0
	for x != 0 {
		parity += PARITY_LOOKUP_TABLE[x&0b11111111]
		x >>= 8
	}
	return parity % 2
}

// using ancient god magic
func parityXor(x int64) int {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return int(x & 0b1)
}
