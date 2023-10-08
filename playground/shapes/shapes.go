package shapes

import "math"

type Rectangle struct{
	Height float64
	Width float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct{
	Radious float64
}

func (c *Circle) Area() float64 {
	return math.Pi * (c.Radious*c.Radious)
}

type Triangle struct{
	Base float64
	Height float64
}

func (r *Triangle) Area() float64 {
	return (r.Base * r.Height) * 0.5
}

func CalcArea(s Shape) float64 {
	return s.Area()
}