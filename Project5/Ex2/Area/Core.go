package Area

import "math"

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
