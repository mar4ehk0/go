package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanBinarySearch(t *testing.T) {
	tests := []struct {
		name          string
		haystack      []int
		needle        int
		expectedIndex int
	}{
		{
			"Items 1, searched value 1",
			[]int{1},
			1,
			0,
		},
		{
			"Items 10, searched first element",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1,
			0,
		},
		{
			"Items 10, searched last element",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
			9,
		},
		{
			"Items 10 with almost equals values",
			[]int{11, 12, 333, 334, 500, 6000, 7999, 8000, 9000, 10000},
			7999,
			6,
		},
		{
			"Items 4s with almost equals values",
			[]int{10, 20, 30, 40, 50, 6000000000000, 7000000000000, 8000000000000, 9000000000000, 10000000000000},
			50,
			4,
		},
		{
			"Items 10 search value 0",
			[]int{-10, -7, -2, -1, 0, 1, 3, 5, 9, 10},
			0,
			4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := BinarySearch(tc.haystack, tc.needle)
			assert.Equal(t, tc.expectedIndex, actual)
		})
	}
}

func TestFailBinarySearchWhenEmpty(t *testing.T) {
	haystack := []int{}
	needle := 1

	_, err := BinarySearch(haystack, needle)

	assert.ErrorIs(t, ErrEmpty, err)
}

func TestFailBinarySearchWhenOneItemAndNotFound(t *testing.T) {
	haystack := []int{1}
	needle := 2

	_, err := BinarySearch(haystack, needle)

	assert.ErrorIs(t, ErrNotFound, err)
}

func TestFailBinarySearchWhenNotFound(t *testing.T) {
	tests := []struct {
		name        string
		haystack    []int
		needle      int
		expectedErr error
	}{
		{
			"Search items in middle closely",
			[]int{1, 2, 4, 5, 6, 7},
			3,
			ErrNotFound,
		},
		{
			"Search items in middle",
			[]int{1, 2, 40000, 50000, 60000, 70000},
			3,
			ErrNotFound,
		},
		{
			"Search items left closely",
			[]int{1, 2, 4, 5, 6, 7},
			0,
			ErrNotFound,
		},
		{
			"Search items left closely",
			[]int{1, 2, 4, 5, 6, 7},
			-10000,
			ErrNotFound,
		},
		{
			"Search items right closely",
			[]int{1, 2, 4, 5, 6, 7},
			8,
			ErrNotFound,
		},
		{
			"Search items right",
			[]int{1, 2, 4, 5, 6, 7},
			80000000,
			ErrNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := BinarySearch(tc.haystack, tc.needle)
			assert.ErrorIs(t, tc.expectedErr, err)
		})
	}
}

func TestFailBinarySearchWHenNotOrdered(t *testing.T) {
	haystack := []int{1, 2, 3, 10, 5, 6, 7, 8}
	needle := 8

	_, err := BinarySearch(haystack, needle)

	assert.ErrorIs(t, ErrNotFound, err)
}
