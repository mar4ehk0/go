package shape

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle(r float64) *Circle {
	circle := &Circle{radius: r}

	return circle
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}

func (c *Circle) HumanDetails() string {
	return fmt.Sprintf("Круг: радиус %f", c.radius)
}
