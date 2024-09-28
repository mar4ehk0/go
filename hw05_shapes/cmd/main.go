package main

import (
	"errors"
	"fmt"

	"github.com/mar4ehk0/go/hw05_shapes/internal/animal"
	"github.com/mar4ehk0/go/hw05_shapes/internal/shape"
)

func main() {
	var someShape [4]any

	someShape[0] = shape.NewCircle(5.0)
	someShape[1] = shape.NewRectangle(10.0, 5.0)
	someShape[2] = shape.NewTriangle(8.0, 6.0)
	someShape[3] = animal.NewDog("Fido")

	for i := 0; i < len(someShape); i++ {
		value, err := calculateArea(someShape[i])
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err.Error())
			continue
		}
		printDetails(someShape[i])
		fmt.Println(value)
	}
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(shape.Shape)
	if ok {
		return shape.Area(), nil
	}

	return 0.0, errors.New("переданный объект не является фигурой")
}

func printDetails(s any) {
	shape, ok := s.(shape.Shape)
	if ok {
		fmt.Println(shape.HumanDetails())
	}
}
