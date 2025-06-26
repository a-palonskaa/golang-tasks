package areacalc

import (
	"math"
	"strings"
)

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

//-------------- Rectangle --------------

type Rectangle struct {
	a    float64
	b    float64
	kind string
}

func NewRectangle(a float64, b float64, str string) *Rectangle {
	return &Rectangle{a, b, str}
}

func (rect *Rectangle) Area() float64 {
	return rect.a * rect.b
}

func (rect *Rectangle) Type() string {
	return rect.kind
}

//-------------- Circle --------------

type Circle struct {
	r    float64
	kind string
}

func NewCircle(r float64, str string) *Circle {
	return &Circle{r, str}
}

func (circle *Circle) Area() float64 {
	return pi * circle.r * circle.r
}

func (circle *Circle) Type() string {
	return circle.kind
}

//-------------- AreaCalculator --------------

func AreaCalculator(figures []Shape) (string, float64) {
	area := float64(0)
	parts := make([]string, 0, len(figures))
	for _, val := range figures {
		parts = append(parts, val.Type())
		area += val.Area()
	}
	return strings.Join(parts, "-"), area
}
