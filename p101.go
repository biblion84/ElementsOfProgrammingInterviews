package main

import "fmt"

func main() {

	fmt.Println(phoneMnemonics([]int{2, 3, 4, 7, 7, 7}))
	fmt.Println(phoneMnemonics([]int{2, 3, 4}))

}

var keypad = map[int][]string{
	1: {},
	2: {"A", "B", "C"},
	3: {"D", "E", "F"},
	4: {"G", "H", "I"},
	5: {"J", "K", "L"},
	6: {"M", "N", "O"},
	7: {"P", "Q", "R", "S"},
	8: {"T", "U", "V"},
	9: {"W", "X", "Y", "Z"},
	0: {},
}

// The variant without recursion
// basically we treat the combination like a number
// but because it's not a fixed base (some digit have 3 values, other 4)
// we add and carry manually for each digits
func phoneMnemonics(in []int) []string {

	ranges := [][]string{}

	for _, c := range in {
		ranges = append(ranges, keypad[c])
	}

	maxCombinations := 1
	for _, r := range ranges {
		maxCombinations *= len(r)
	}

	rangeKey := make([]int, len(ranges))
	combinations := make([]string, 0, maxCombinations)
	for i := 0; i < maxCombinations; i++ {
		result := ""
		for j, k := range rangeKey {
			result += ranges[j][k]
		}

		combinations = append(combinations, result)

		for j := len(rangeKey) - 1; j >= 0; j-- {
			rangeKey[j]++
			if rangeKey[j] >= len(ranges[j]) {
				rangeKey[j] = 0
			} else {
				break
			}
		}
	}
	return combinations
}
