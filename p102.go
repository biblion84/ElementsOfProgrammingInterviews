package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(lookAndSay(1))
	fmt.Println(lookAndSay(11))
	fmt.Println(lookAndSay(21))
	fmt.Println(lookAndSay(1211))
	fmt.Println(lookAndSay(111221))
	fmt.Println(lookAndSay(312211))

}

func lookAndSay(in int) string {

	number := []rune(strconv.Itoa(in))

	result := ""

	previous := number[0]
	count := 1
	for _, d := range number[1:] {
		if d != previous {
			result += fmt.Sprintf("%d%c", count, previous)
			count = 0
		}
		previous = d
		count++
	}
	result += fmt.Sprintf("%d%c", count, previous)

	return result

}
