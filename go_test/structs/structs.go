package structs

import "math"

type Rectagle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectagle) float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectagle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
