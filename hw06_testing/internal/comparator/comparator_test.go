package comparator

import (
	"testing"

	"github.com/mar4ehk0/go/hw06_testing/internal/comparator/book"
	"github.com/mar4ehk0/go/hw06_testing/internal/comparator/checker"
	"github.com/stretchr/testify/assert"
)

func TestCanCheck(t *testing.T) {
	// arrange
	tests := []struct {
		name           string
		bookA          *book.Book
		bookB          *book.Book
		comparatorMode ModeEnum
		expected       string
	}{
		{
			"Year: A > B",
			book.CreateBook(1, "BookA", "Author Book A", 1901, 100, 5),
			book.CreateBook(2, "BookB", "Author Book B", 1901, 100, 5),
			Year,
			checker.CheckBigger,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// act
			dto := checker.NewDto(
				tc.comparatorMode,
				*tc.bookA,
				*tc.bookB,
			)
			actual := checker.Check(dto)
			// assert
			assert.Equal(t, tc.expected, actual)
		})
	}
}
