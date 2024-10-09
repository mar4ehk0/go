package search

const NotFound = -1

func BinarySearch(haystack []int, needle int) int {
	var guess int

	size := len(haystack)

	if size < 1 {
		return NotFound
	}

	if size == 1 {
		if haystack[0] == needle {
			return 0
		}

		return NotFound
	}

	left := 0
	right := size - 1

	if !(haystack[left] <= needle && needle <= haystack[right]) {
		return NotFound
	}

	for left <= right {
		mid := left + (right-left)/2
		guess = haystack[mid]
		if guess == needle {
			return mid
		}
		if needle > guess {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return NotFound
}
