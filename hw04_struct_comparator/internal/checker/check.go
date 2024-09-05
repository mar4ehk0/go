package checker

import (
	"fmt"

	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/book"
	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/comparator"
)

type Dto struct {
	Mode  comparator.ModeEnum
	BookA book.Book
	BookB book.Book
}

func Check(collection ...Dto) {
	for _, dto := range collection {
		cmp := comparator.NewComparator(dto.Mode)
		bookA := dto.BookA
		bookB := dto.BookB

		result := cmp.Compare(bookA, bookB)

		var value string

		if result {
			value = "bigger"
		} else {
			value = "smaller"
		}

		fmt.Printf("%s of book: \"%s\" %s then book: \"%s\" \n", dto.Mode, bookA.GetTitle(), value, bookB.GetTitle())
	}
}
