package main

import "fmt"

type Rectangle struct {
	Xmin, Xmax int
	Ymin, Ymax int
}

func main() {

	fmt.Println(intersectingRectangle(Rectangle{
		Xmin: 0, Xmax: 2,
		Ymin: 0, Ymax: 2,
	}, Rectangle{
		Xmin: 2, Xmax: 4,
		Ymin: 2, Ymax: 4,
	}))

	fmt.Println(intersectingRectangle(Rectangle{
		Xmin: 0, Xmax: 3,
		Ymin: 1, Ymax: 2,
	}, Rectangle{
		Xmin: 1, Xmax: 2,
		Ymin: 0, Ymax: 3,
	}))

}

func intersectingRectangle(a, b Rectangle) Rectangle {
	if !intersect(a, b) {
		return Rectangle{}
	}

	intersection := Rectangle{}

	intersection.Xmin = max(a.Xmin, b.Xmin)
	intersection.Ymin = max(a.Ymin, b.Ymin)
	intersection.Xmax = min(a.Xmax, b.Xmax)
	intersection.Ymax = min(a.Ymax, b.Ymax)

	return intersection

}

func intersect(a, b Rectangle) bool {
	// we are in a x boundary
	intersectX := a.Xmin <= b.Xmax && a.Xmax >= b.Xmin
	intersectY := a.Ymin <= b.Ymax && a.Ymax >= b.Ymin
	return intersectX && intersectY
}
