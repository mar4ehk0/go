package comparator

import (
	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/book"
)

type Comparator struct {
	mode ModeEnum
}

func NewComparator(mode ModeEnum) Comparator {
	comparator := Comparator{mode}

	return comparator
}

func (c Comparator) Compare(a book.Book, b book.Book) bool {
	switch c.mode {
	case Year:
		return a.Year() >= b.Year()
	case Size:
		return a.Size() >= b.Size()
	case Rate:
		return a.Rate() >= b.Rate()
	default:
		return false
	}
}
