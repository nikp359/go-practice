package structs

type Rectagle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectagle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(r Rectagle) float64 {
	return r.Width * r.Height
}
