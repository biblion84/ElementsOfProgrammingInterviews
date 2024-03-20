package main

import "fmt"

func main() {
	x := int8(5)
	y := int8(21)
	result := int8(0)

	for y > 0 {
		result = add(result, x)
		y = add(y, -1)
	}

	fmt.Println(result)
	//fmt.Println(add(5, -1))
}
func add(a, b int8) int8 {
	fmt.Printf("a: %s\nb: %s\n------\n", pb(a), pb(b))
	for b != 0 {
		carry := a & b
		a ^= b
		b = carry << 1
		fmt.Printf("c: %s\na: %s\nb: %s\n---------\n", pb(carry), pb(a), pb(b))
	}
	fmt.Printf("result: %d\n", a)
	return a
}

// printBinary
func pb(in int8) string {
	result := fmt.Sprintf(" (%d)", in)
	for i := 0; i < 8; i++ {
		if (in & 1) == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}
		in >>= 1
	}
	return result
}

// the XOR is doing an ADD without the carry
// 1 + 1 = 0 (missing carry)
// 1 + 0 = 1
// 0 + 0 = 0
// the AND is doing only the carry
// 1 + 1 = 1
// 1 + 0 = 0
// 0 + 0 = 0
// the SHIFT LEFT is propagating the carry
// we repeat by adding, propagating the carry until no carry is left

// add(5,3)
//	a: 00000101 (5)
//	b: 00000011 (3)
//	------
//	c: 00000001 (1)
//	a: 00000110 (6)
//	b: 00000010 (2)
//	---------
//	c: 00000010 (2)
//	a: 00000100 (4)
//	b: 00000100 (4)
//	---------
//	c: 00000100 (4)
//	a: 00000000 (0)
//	b: 00001000 (8)
//	---------
//	c: 00000000 (0)
//	a: 00001000 (8)
//	b: 00000000 (0)
//	---------
//	8

// thanks to 2 complement, it also work when adding a negative number to a positive one
//	add(5,-1)
//	a: 00000101 (5)
//	b: 11111111 (-1)
//	------
//	c: 00000101 (5)
//	a: 11111010 (-6)
//	b: 00001010 (10)
//	---------
//	c: 00001010 (10)
//	a: 11110000 (-16)
//	b: 00010100 (20)
//	---------
//	c: 00010000 (16)
//	a: 11100100 (-28)
//	b: 00100000 (32)
//	---------
//	c: 00100000 (32)
//	a: 11000100 (-60)
//	b: 01000000 (64)
//	---------
//	c: 01000000 (64)
//	a: 10000100 (-124)
//	b: 10000000 (-128)
//	---------
//	c: 10000000 (-128)
//	a: 00000100 (4)
//	b: 00000000 (0)
//	---------
//	4
