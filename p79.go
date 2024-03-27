package main

import (
	"fmt"
	"math/rand"
)

func main() {

	size := 5
	a := make([]int, size)

	for i := 0; i < size; i++ {
		if i > (size / 20) {
			a[i] = 0
		} else {
			a[i] = 1
		}
	}

	fmt.Println(sniffer(a, 10))
	fmt.Println(sniffer(a, 10))
	fmt.Println(sniffer(a, 10))
	fmt.Println(sniffer(a, 10))
	fmt.Println(sniffer(a, 10))
	fmt.Println(sniffer(a, 10))

}

// in should be treated as a stream of unknown size
func sniffer(in []int, size int) []int {
	if len(in) <= size {
		return in
	}

	buffer := make([]int, size)
	seen := 0

	// fill the buffer with the size first size elements
	for i := 0; i < len(in) && i < len(buffer); i++ {
		buffer[i] = in[i]
	}
	in = in[size+1:]

	for i := 0; i < len(in); i++ {
		seen++
		// for each element, we add it to our sniffer buffer with a decreasing probability
		sample := rand.Intn(seen) < size
		if sample {
			buffer[rand.Intn(size)] = in[i]
		}
	}

	return buffer
}
