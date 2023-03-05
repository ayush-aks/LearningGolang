package main

import (
	"fmt"
)

// Write a program that creates two custom struct types called 'triangle' and 'square' and interface shape

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{sideLength: 4.0}
	tr := triangle{height: 2.0, base: 4.0}
	printArea(sq)
	printArea(tr)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

