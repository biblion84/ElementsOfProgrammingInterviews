package main

import "fmt"

// 5.4
func main() {
	fmt.Println(closest(21))
	fmt.Println(closest(6))

}

func closest(x int64) int64 {
	// we need to swap the two rightmost bits that differs

	for i := 0; i < 63; i++ {
		if ((x >> i) & 1) != ((x >> (i + 1)) & 1) {
			mask := int64(1<<i | 1<<(i+1))
			x ^= mask
			return x
		}
	}

	return 0
}
