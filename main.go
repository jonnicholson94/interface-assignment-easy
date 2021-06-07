package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	width  float64
}

func main() {
	sq := square{sideLength: 8.00}
	t := triangle{height: 4.00, width: 5.00}

	printArea(sq)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.width
}
