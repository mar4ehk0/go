package search

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrEmpty    = errors.New("empty")
)

func BinarySearch(haystack []int, needle int) (int, error) {
	var guess int

	size := len(haystack)

	if size < 1 {
		return 0, ErrEmpty
	}

	if size == 1 {
		if haystack[0] == needle {
			return 0, nil
		}

		return 0, ErrNotFound
	}

	left := 0
	right := size - 1

	for left <= right {
		mid := left + (right-left)/2
		guess = haystack[mid]
		if guess == needle {
			return mid, nil
		}
		if needle > guess {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0, ErrNotFound
}
