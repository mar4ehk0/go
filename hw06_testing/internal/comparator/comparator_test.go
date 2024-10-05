package comparator

import (
	"testing"

	"github.com/mar4ehk0/go/hw06_testing/internal/comparator/book"
	"github.com/stretchr/testify/assert"
)

func TestCanCheck(t *testing.T) {
	tests := []struct {
		name           string
		bookA          *book.Book
		bookB          *book.Book
		comparatorMode ModeEnum
		expected       int8
	}{
		/////////Year//////////////
		{
			"Year: A > B",
			book.CreateBook(1, "BookA", "Author Book A", 10, 100, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1, 100, 5),
			ComparatorModeYear,
			ComparatorBLess,
		},
		{
			"Year: A < B",
			book.CreateBook(1, "BookA", "Author Book A", 0, 100, 5),
			book.CreateBook(2, "BookB", "Author Book B", 3, 100, 5),
			ComparatorModeYear,
			ComparatorALess,
		},
		{
			"Year: A == B",
			book.CreateBook(1, "BookA", "Author Book A", 1, 100, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1, 100, 5),
			ComparatorModeYear,
			ComparatorEqual,
		},
		/////////Size//////////////
		{
			"Size: A > B",
			book.CreateBook(1, "BookA", "Author Book A", 10, 100, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1, 10, 5),
			ComparatorModeYear,
			ComparatorBLess,
		},
		{
			"Size: A < B",
			book.CreateBook(1, "BookA", "Author Book A", 0, 100, 5),
			book.CreateBook(2, "BookB", "Author Book B", 3, 9999, 5),
			ComparatorModeYear,
			ComparatorALess,
		},
		{
			"Size: A == B",
			book.CreateBook(1, "BookA", "Author Book A", 1, 99999, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1, 99999, 5),
			ComparatorModeYear,
			ComparatorEqual,
		},
		{
			"Size: A == B, huge number",
			book.CreateBook(1, "BookA", "Author Book A", 1, 9999999999999999999, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1, 9999999999999999999, 5),
			ComparatorModeYear,
			ComparatorEqual,
		},
		/////////Rate//////////////
		{
			"Rate: A > B",
			book.CreateBook(1, "BookA", "Author Book A", 10, 100, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1, 10, 2.333),
			ComparatorModeYear,
			ComparatorBLess,
		},
		{
			"Rate: A < B",
			book.CreateBook(1, "BookA", "Author Book A", 0, 100, 1),
			book.CreateBook(2, "BookB", "Author Book B", 3, 9999, 2.5),
			ComparatorModeYear,
			ComparatorALess,
		},
		{
			"Rate: A == B",
			book.CreateBook(1, "BookA", "Author Book A", 1, 99999, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1, 99999, 5),
			ComparatorModeYear,
			ComparatorEqual,
		},
		{
			"Rate: A < B, almost equal",
			book.CreateBook(1, "BookA", "Author Book A", 1, 99999, 4.9999999999999999999),
			book.CreateBook(2, "BookB", "Author Book B", 1, 99999, 5),
			ComparatorModeYear,
			ComparatorEqual,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := Check(tc.comparatorMode, tc.bookA, tc.bookB)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestFailCheck(t *testing.T) {
	tests := []struct {
		name           string
		bookA          *book.Book
		bookB          *book.Book
		comparatorMode ModeEnum
		expected       error
	}{
		{
			"Unknown mode",
			book.CreateBook(1, "BookA", "Author Book A", 10, 100, 5),
			book.CreateBook(1, "BookA", "Author Book A", 10, 100, 5),
			123,
			ErrUnknownModeEnum,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := Check(tc.comparatorMode, tc.bookA, tc.bookB)
			assert.ErrorIs(t, tc.expected, err)
		})
	}
}
