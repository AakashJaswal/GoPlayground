package main

import "fmt"

func main() {
	sq := square{sidelength: 10}
	tr := triangle{height: 10, base: 10}

	printArea(sq)

	printArea(tr)

}

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

func printArea(s shape) {
	println(fmt.Sprint(s.getArea()))
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}
