package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	// TODO: build a Point for each pair of coordinates, then print the squared distance between them instead of this placeholder.
	p1 := Point{x1, y1}
	p2 := Point{x2, y2}

	dx := p1.X - p2.X
	dy := p1.Y - p2.Y

	squaredDistance := dx*dx + dy*dy

	fmt.Println(squaredDistance)
}
