package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Bob likes Alice"))
}

func reverseWords(in string) string {
	split := strings.Split(in, " ")
	slices.Reverse(split)
	return strings.Join(split, " ")
}
