package checker

import (
	"fmt"

	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/book"
	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/comparator"
)

type Dto struct {
	mode  comparator.ModeEnum
	bookA book.Book
	bookB book.Book
}

func NewDto(mode comparator.ModeEnum, bookA book.Book, bookB book.Book) Dto {
	return Dto{mode: mode, bookA: bookA, bookB: bookB}
}

func Check(dto Dto) {
	cmp := comparator.NewComparator(dto.mode)
	bookA := dto.bookA
	bookB := dto.bookB

	result := cmp.Compare(bookA, bookB)

	var value string

	if result {
		value = "bigger"
	} else {
		value = "smaller"
	}

	fmt.Printf("%s of book: \"%s\" %s then book: \"%s\" \n", dto.mode, bookA.Title(), value, bookB.Title())
}
