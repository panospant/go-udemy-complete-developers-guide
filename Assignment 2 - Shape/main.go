package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sidelength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {

	triangle := triangle{height: 1, base: 7}
	square := square{sidelength: 2}

	printArea(triangle)
	printArea(square)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}
