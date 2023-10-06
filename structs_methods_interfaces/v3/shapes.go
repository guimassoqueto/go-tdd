package main

import "math"

type Shape interface {
	Area() float64
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}


type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.height + rectangle.width)
}


type Triangle struct {
	base float64
	height float64
}

func (triangle Triangle) Area() float64 {
	return (triangle.base * triangle.height) * 0.5
}