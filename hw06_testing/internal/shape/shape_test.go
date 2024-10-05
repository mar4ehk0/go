package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCalculateArea(t *testing.T) {
	// arrange
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			"Circle",
			NewCircle(5.0),
			78.53981633974483,
		},
		{
			"Rectangle",
			NewRectangle(10.0, 5.0),
			50,
		},
		{
			"Triangle",
			NewTriangle(8.0, 6.0),
			24,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// act
			actual, _ := CalculateArea(tc.shape)

			// assert
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestFailWhenNotShape(t *testing.T) {
	tests := []struct {
		name     string
		shape    any
		expected error
	}{
		{
			"Int",
			10,
			ErrNotShape,
		},
		{
			"0",
			0,
			ErrNotShape,
		},
		{
			"-1",
			-1,
			ErrNotShape,
		},
		{
			"nil",
			nil,
			ErrNotShape,
		},
		{
			"Custom Struct",
			struct{ userName string }{userName: "Lorem Ipsum"},
			ErrNotShape,
		},
		{
			"String",
			"LoremIpsum",
			ErrNotShape,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := CalculateArea(tc.shape)

			assert.ErrorIs(t, tc.expected, err)
		})
	}
}
