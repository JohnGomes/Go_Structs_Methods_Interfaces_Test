package main

import "math"

type Shapes interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

//--------------------------------------------------------

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//--------------------------------------------------------

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
