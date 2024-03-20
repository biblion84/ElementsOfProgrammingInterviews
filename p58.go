package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// X, Y aligned rectangle
	fmt.Println(isRectangleNotAligned([4]Point{
		{0, 0},
		{1, 0},
		{1, 2},
		{0, 2},
	}))
	// X, Y unaligned rectangle
	fmt.Println(isRectangleNotAligned([4]Point{
		{0, 3},
		{2, 5},
		{3, 0},
		{5, 2},
	}))
	// X, Y unaligned non-rectangle
	fmt.Println(isRectangleNotAligned([4]Point{
		{0, 3},
		{2, 5},
		{3, 0},
		{5, 3},
	}))
	// X, Y aligned square
	fmt.Println(isRectangleNotAligned([4]Point{
		{0, 2},
		{2, 2},
		{0, 0},
		{2, 0},
	}))
	// rhombus // diamond
	fmt.Println(isRectangleNotAligned([4]Point{
		{0, 2},
		{1, 4},
		{1, 0},
		{2, 2},
	}))
}

// Note: this assume that the rectangle is X and Y aligned
// given 4 points in the plane, how would you check if they are the vertices of a rectangle ?
func isRectangle(points [4]Point) bool {

	x := map[int]int{}
	y := map[int]int{}

	for _, p := range points {
		x[p.X]++
		y[p.Y]++
	}

	// to be a rectangle we must have 2 x and 2 y that's all
	for _, occurence := range x {
		if occurence != 2 {
			return false
		}
	}
	for _, occurence := range y {
		if occurence != 2 {
			return false
		}
	}

	return true

}

func squared(a int) int {
	return a * a
}

func distance(p1, p2 Point) int {
	return squared(p1.X-p2.X) + squared(p1.Y-p2.Y)
}

// This allows to find is the 4 points are vertices of a rectangle even when it is not X, Y aligned
func isRectangleNotAligned(points [4]Point) bool {
	distances := map[int]int{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distances[distance(points[i], points[j])]++
		}
	}

	// in a rectangle we have 3 max different distances : Width, Length, Diagonal
	if len(distances) != 2 && len(distances) != 3 {
		return false
	}

	// if it's a rectangle we should have either
	// 2 different values (4 width and 2 diagonal) if we're a square
	// or 3 different values:  2 width, 2 length, 2 diagonals
	if len(distances) == 2 {
		for _, occurence := range distances {
			if occurence != 4 && occurence != 2 {
				return false
			}
		}
		return true
	}

	if len(distances) == 3 {
		for _, occurence := range distances {
			if occurence != 2 {
				return false
			}
		}
		return true
	}

	return false

}
