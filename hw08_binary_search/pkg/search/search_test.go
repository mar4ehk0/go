package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		haystack []int
		needle   int
		expected int
	}{
		// success
		{
			"Found: Items 1, searched value 1",
			[]int{1},
			1,
			0,
		},
		{
			"Found: Items 10, searched first element",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1,
			0,
		},
		{
			"Found: Items 10, searched last element",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
			9,
		},
		{
			"Found: Items 10 with almost equals values",
			[]int{11, 12, 333, 334, 500, 6000, 7999, 8000, 9000, 10000},
			7999,
			6,
		},
		{
			"Found: Items 4s with almost equals values",
			[]int{10, 20, 30, 40, 50, 6000000000000, 7000000000000, 8000000000000, 9000000000000, 10000000000000},
			50,
			4,
		},
		{
			"Found: Items 10 search value 0",
			[]int{-10, -7, -2, -1, 0, 1, 3, 5, 9, 10},
			0,
			4,
		},
		// not found
		{
			"Not Found: Items empty, not found",
			[]int{},
			1,
			NotFound,
		},
		{
			"Not found: Items contains one item, not found",
			[]int{1},
			2,
			NotFound,
		},
		{
			"Not Found: Search items in middle closely",
			[]int{1, 2, 4, 5, 6, 7},
			3,
			NotFound,
		},
		{
			"Not Found: Search items in middle",
			[]int{1, 2, 40000, 50000, 60000, 70000},
			3,
			NotFound,
		},
		{
			"Not Found: Search items left closely",
			[]int{1, 2, 4, 5, 6, 7},
			0,
			NotFound,
		},
		{
			"Not Found: Search items left",
			[]int{1, 2, 4, 5, 6, 7},
			-10000,
			NotFound,
		},
		{
			"Not Found: Search items right closely",
			[]int{1, 2, 4, 5, 6, 7},
			8,
			NotFound,
		},
		{
			"Not Found: Search items right",
			[]int{1, 2, 4, 5, 6, 7},
			80000000,
			NotFound,
		},
		{
			"Not Found: Search not ordered items",
			[]int{1, 2, 3, 10, 5, 6, 7, 8},
			8,
			NotFound,
		},
		{
			"Not Found: Search for out of bounds value, bigger",
			[]int{1, 2, 3, 10},
			888,
			NotFound,
		},
		{
			"Not Found: Search for out of bounds value, less",
			[]int{1, 2, 3, 10},
			-456,
			NotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := BinarySearch(tc.haystack, tc.needle)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
