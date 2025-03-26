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
	base   float64
}

func main() {
	square := square{sideLength: 5}
	triangle := triangle{height: 5, base: 3}

	printArea(square)
	printArea(triangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	side := s.sideLength
	area := side * side
	return area
}

func (t triangle) getArea() float64 {
	// could return without saving the data as variables
	base := t.base
	height := t.height
	area := base * height * 0.5
	return area
}
