package checker

import (
	"github.com/mar4ehk0/go/hw06_testing/internal/comparator"
	"github.com/mar4ehk0/go/hw06_testing/internal/comparator/book"
)

const (
	CheckBigger  string = "bigger"
	CheckSmaller string = "smaller"
)

type Dto struct {
	mode  comparator.ModeEnum
	bookA book.Book
	bookB book.Book
}

func NewDto(mode comparator.ModeEnum, bookA book.Book, bookB book.Book) *Dto {
	return &Dto{mode: mode, bookA: bookA, bookB: bookB}
}

func Check(dto *Dto) string {
	cmp := comparator.NewComparator(dto.mode)
	bookA := dto.bookA
	bookB := dto.bookB

	result := cmp.Compare(bookA, bookB)

	var value string

	if result {
		value = CheckBigger
	} else {
		value = CheckSmaller
	}

	return value
}
