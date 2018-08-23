package shapes

import (
	"math"
)

// Shapes calculates Area
type Shape interface {
	Area() float64
}

// Perimeter returns perimeter of a rectangle
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Area returns area of a rectangle
func Area(width float64, height float64) float64 {
	return width * height
}

// Rectangle defines parameters of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area return area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle defines parameters of a Circle
type Circle struct {
	radius float64
}

// Area return area of the Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Triangle defines parameters of a Triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area return area of the Triangle
func (t Triangle) Area() float64 {
	return t.Height * t.Base * 0.5
}
