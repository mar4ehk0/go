package shape

import (
	"fmt"
)

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base, height float64) *Triangle {
	triangle := &Triangle{base: base, height: height}

	return triangle
}

func (t *Triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t *Triangle) HumanDetails() string {
	return fmt.Sprintf("Треугольник: основание %f, высота %f", t.base, t.height)
}
