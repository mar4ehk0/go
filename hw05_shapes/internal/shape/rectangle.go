package shape

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	rectangle := &Rectangle{width: width, height: height}

	return rectangle
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) HumanDetails() string {
	return fmt.Sprintf("Прямоугольник: ширина %f, высота %f", r.width, r.height)
}
