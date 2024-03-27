package main

import "fmt"

func main() {

	fmt.Println(itoaBase(atoiBase("FF", 16), 2))
	fmt.Println(atoiBase("101", 2))
	fmt.Println(atoiBase("FF", 16))

}

// works from base 2 to base 36 because we only use [0-9] + [A-Z] as digits
func atoiBase(in string, base int) int {
	negative := in[0] == '-'
	if negative {
		in = in[1:]
	}

	result := 0

	for _, c := range in {

		value := int(c - '0')
		if value < 0 || value > 9 {
			value = int(c-'A') + 10
			if value < 10 || value > base {
				panic("bad input")
			}
		}
		result *= base
		result += value
	}

	if negative {
		result *= -1
	}

	return result
}

func itoaBase(in int, base int) string {
	if in == 0 {
		return "0"
	}
	result := ""

	for in > 0 {
		value := in % base
		if value < 10 {
			result += string('0' + value)
		} else {
			result += string('A' + value - 10)
		}
		in /= base
	}

	return result
}
