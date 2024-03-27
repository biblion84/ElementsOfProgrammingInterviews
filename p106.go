package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(runLengthEncoding("aaaabcccaa"))
	fmt.Println(runLengthDecoding("3r4b"))
	fmt.Println(runLengthDecoding(runLengthEncoding("aaaabcccaa")))

}

// assuming no digits in the input
func runLengthEncoding(in string) string {
	r := []rune(in)

	result := ""

	previous := r[0]
	count := 1
	for _, c := range r[1:] {
		if c != previous {
			result += fmt.Sprintf("%d%c", count, previous)
			previous = c
			count = 0
		}
		count++
	}
	result += fmt.Sprintf("%d%c", count, previous)

	return result
}

func runLengthDecoding(in string) string {
	result := ""

	count := ""
	for _, c := range in {
		if c >= '0' && c <= '9' {
			count += string(c)
		} else {
			countNumber, err := strconv.Atoi(count)
			count = ""
			if err != nil {
				panic("wrongly formatted input")
			}
			for i := 0; i < countNumber; i++ {
				result += string(c)
			}
		}
	}

	return result
}
