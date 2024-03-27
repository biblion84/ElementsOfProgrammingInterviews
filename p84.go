package main

import "fmt"

func main() {

	// Thanks chatgpt for the sudoku encoding, love you
	sudoku := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println(sudoku[:3])

	fmt.Println(sudoku[:][1])

	//fmt.Println(isSudokuValid(sudoku))

}

func subregionValid(sudoku [][]int, xMin, xMax, yMin, yMax int) bool {
	found := [10]bool{}

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			value := sudoku[y][x]
			if value == 0 {
				continue
			}
			if found[value] {
				return false
			}
			found[value] = true
		}
	}

	return true
}

func isSudokuValid(sudoku [][]int) bool {
	assert(len(sudoku) == 9 && len(sudoku[0]) == 9)

	for x := 0; x < len(sudoku[0]); x++ {
		if !subregionValid(sudoku, x, x, 0, len(sudoku)-1) {
			return false
		}
	}

	for y := 0; y < len(sudoku); y++ {
		if !subregionValid(sudoku, 0, len(sudoku[y])-1, y, y) {
			return false
		}
	}

	for y := 0; y < len(sudoku); y += 3 {
		for x := 0; x < len(sudoku[0]); x += 3 {
			if !subregionValid(sudoku, x, x+2, y, y+2) {
				return false
			}
		}
	}

	return true

}

func assert(cond bool) {
	if !cond {
		panic("")
	}
}
