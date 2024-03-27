package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	//
	fmt.Println(itoa(atoi("9893898")))
	fmt.Println(itoa(atoi("-9893898")))
	fmt.Println(itoa(atoi("0")))
}

func atoi(in string) int {
	negative := in[0] == '-'
	if negative {
		in = in[1:]
	}
	result := 0

	for _, c := range in {
		value := int(c - '0')
		if value < 0 || value > 9 {
			panic("")
		}
		result *= 10
		result += value
	}

	if negative {
		result *= -1
	}

	return result
}

func itoa(in int) string {
	if in == 0 {
		return "0"
	}

	negative := in < 0
	length := 0
	if negative {
		in *= -1
		length++
	}

	// The number of digits of a number is the floor of its log + 1
	length += int(math.Floor(math.Log10(float64(in)))) + 1

	result := make([]rune, 0, length)

	for in > 0 {
		result = append(result, rune(in%10+'0'))
		in /= 10
	}

	if negative {
		result = append(result, '-')
	}
	slices.Reverse(result)

	return string(result)

}
