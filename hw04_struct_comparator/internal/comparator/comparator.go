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
		return a.GetYear() >= b.GetYear()
	case Size:
		return a.GetSize() >= b.GetSize()
	case Rate:
		return a.GetRate() >= b.GetRate()
	default:
		return false
	}
}
