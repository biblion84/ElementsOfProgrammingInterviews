package main

import "fmt"

func main() {

	fmt.Println(excelEncoding("A"))
	fmt.Println(excelEncoding("B"))
	fmt.Println(excelEncoding("AA"))
	fmt.Println(excelEncoding("ZZ"))
}

func excelEncoding(in string) int {
	result := 0

	for _, c := range in {
		result *= 26
		result += int(c - 'A' + 1)
	}

	return result
}
