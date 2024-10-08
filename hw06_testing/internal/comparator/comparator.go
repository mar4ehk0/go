package comparator

import "github.com/mar4ehk0/go/hw06_testing/internal/comparator/book"

const (
	ComparatorEqual int8 = 0
	ComparatorALess int8 = -1
	ComparatorBLess int8 = 1
)

func Check(mode ModeEnum, bookA *book.Book, bookB *book.Book) (int8, error) {
	var result int8
	var err error

	switch mode {
	case ComparatorModeYear:
		result = compareInt(bookA.Year(), bookB.Year())
	case ComparatorModeSize:
		result = compareInt(bookA.Size(), bookB.Size())
	case ComparatorModeRate:
		result = compareFloat(bookA.Rate(), bookB.Rate())
	default:
		err = ErrUnknownModeEnum
	}

	return result, err
}

func compareInt(a uint, b uint) int8 {
	if a < b {
		return ComparatorALess
	} else if a > b {
		return ComparatorBLess
	}

	return ComparatorEqual
}

func compareFloat(a float32, b float32) int8 {
	if a < b {
		return ComparatorALess
	} else if a > b {
		return ComparatorBLess
	}

	return ComparatorEqual
}
