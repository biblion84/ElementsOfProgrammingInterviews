package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(spiralOrder(grid))
}

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
	DIRECTION_MEMBERS
)

// {x, y}
var SHIFT = map[int][2]int{
	RIGHT: {1, 0},
	DOWN:  {0, 1},
	LEFT:  {-1, 0},
	UP:    {0, -1}}

func spiralOrder(in [][]int) []int {
	// assuming a square in
	result := make([]int, 0, len(in)*len(in))
	direction := RIGHT

	x, y := 0, 0
	for i := 0; i < len(in)*len(in); i++ {
		result = append(result, in[y][x])
		// assuming -1 never appear in the input
		in[y][x] = -1
		nextX := x + SHIFT[direction][0]
		nextY := y + SHIFT[direction][1]
		if nextX < 0 || nextX >= len(in) || nextY < 0 || nextY >= len(in) || in[nextY][nextX] == -1 {
			direction = (direction + 1) % DIRECTION_MEMBERS
			nextX = x + SHIFT[direction][0]
			nextY = y + SHIFT[direction][1]
		}
		x, y = nextX, nextY

	}

	return result
}
