package shape

import (
	"errors"
)

var ErrNotShape = errors.New("переданный объект не является фигурой")

func CalculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if ok {
		return shape.Area(), nil
	}

	return 0.0, ErrNotShape
}
