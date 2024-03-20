package main

import "fmt"

// 5.2
func main() {

	fmt.Println(swapBit(0b01001001, 1, 6))
	fmt.Println(swapBitMask(0b01001001, 1, 6))

}

func swapBit(x int64, i, j int) int64 {
	iValue := x & (1 << i)
	jValue := x & (1 << j)

	if iValue > 0 {
		x |= 1 << j
	} else {
		x &= ^(1 << j)
	}

	if jValue > 0 {
		x |= 1 << i
	} else {
		x &= ^(1 << i)
	}

	return x
}

func swapBitMask(x int64, i, j int) int64 {
	if (x>>i)&1 != (x>>j)&1 {
		mask := int64(1<<i | 1<<j)
		x ^= mask
	}
	return x
}
