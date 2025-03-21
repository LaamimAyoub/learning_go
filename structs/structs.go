package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	one float64
	two float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.height + rect.width)
}

func (rect Rectangle) Area() float64 {
	return rect.height * rect.width
}

func (circ Circle) Area() float64 {
	return circ.Radius * circ.Radius * math.Pi
}

func (tria Triangle) Area() float64 {
	return tria.one * tria.two
}
