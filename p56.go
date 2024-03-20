package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := 6

	for i := 0; i < 100; i++ {
		fmt.Println(randWithCoin(x))
	}

}

// to generate a number between 0 and 6 using only a coin
// my first intuition would be to construct a number bit by bit (3 bit in this case)
// and for each bit, throw the coin to see if we set it to one or zero
func randWithCoin(x int) int {
	for {
		result := 0
		for i := 0; (1 << i) < x; i++ {
			result <<= 1
			result |= flipCoin() & 1
		}
		if result < x {
			return result
		}
	}
}

func flipCoin() int {
	return rand.Intn(2)
}
