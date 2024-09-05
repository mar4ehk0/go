package main

import (
	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/book"
	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/checker"
	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/comparator"
)

func main() {
	book1 := book.CreateBook(
		11,
		"Lord of rings",
		"JR Tolkien",
		1901,
		600,
		4.59,
	)

	book2 := book.CreateBook(
		22,
		"Война и мир",
		"Л. Н. Толстой",
		1900,
		6000,
		4.0,
	)

	checkerYear := checker.Dto{
		Mode:  comparator.Year,
		BookA: book1,
		BookB: book2,
	}
	checkerSize := checker.Dto{
		Mode:  comparator.Size,
		BookA: book1,
		BookB: book2,
	}
	checkerRate := checker.Dto{
		Mode:  comparator.Rate,
		BookA: book1,
		BookB: book2,
	}

	checker.Check(checkerYear, checkerSize, checkerRate)
}

// func exampleComparatorYear(a book.Book, b book.Book) {
// 	cmp := comparator.NewComparator(comparator.Year)
// 	result := cmp.Compare(a, b)

// 	if result {
// 		fmt.Printf("Year of book: \"%s\" bigger then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
// 	} else {
// 		fmt.Printf("Year of book: \"%s\" smaller then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
// 	}
// }

// func exampleComparatorSize(a book.Book, b book.Book) {
// 	cmp := comparator.NewComparator(comparator.Size)
// 	result := cmp.Compare(a, b)

// 	if result {
// 		fmt.Printf("Size of book: \"%s\" bigger then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
// 	} else {
// 		fmt.Printf("Size of book: \"%s\" smaller then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
// 	}
// }

// func exampleComparatorRate(a book.Book, b book.Book) {
// 	cmp := comparator.NewComparator(comparator.Rate)
// 	result := cmp.Compare(a, b)

// 	if result {
// 		fmt.Printf("Rate of book: \"%s\" bigger then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
// 	} else {
// 		fmt.Printf("Rate of book: \"%s\" smaller then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
// 	}
// }
