package main

import (
	"fmt"
)

func main() {

	fmt.Println(isPalindrome("test"))
	fmt.Println(isPalindrome("kayak"))
	fmt.Println(isPalindrome("A man, a plan, a canal, Panama."))

}

// assume ascii input, the utf8 version is a bit more involved to do in O(1) space
func isPalindrome(in string) bool {
	i, j := 0, len(in)-1

	for i < j {
		for ; !isAlpha(in[i]) && i < j; i++ {
		}

		for ; !isAlpha(in[j]) && i < j; j-- {
		}

		if toLower(in[i]) != toLower(in[j]) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func toLower(c byte) byte {
	if c < 'a' {
		return c + 'a' - 'A'
	}
	return c
}
