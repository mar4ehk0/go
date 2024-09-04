package main

import (
	"fmt"

	"github.com/mar4ehk0/go/hw04_struct_comparator/internal/book"
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

	exampleComparatorYear(book1, book2)
	exampleComparatorSize(book1, book2)
	exampleComparatorRate(book1, book2)
}

func exampleComparatorYear(a book.Book, b book.Book) {
	cmp := comparator.NewComparator(comparator.Year)
	result, err := cmp.Compare(a, b)
	if err != nil {
		fmt.Println("something wrong")
		return
	}

	if result {
		fmt.Printf("Year of book: \"%s\" bigger then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
	} else {
		fmt.Printf("Year of book: \"%s\" smaller then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
	}
}

func exampleComparatorSize(a book.Book, b book.Book) {
	cmp := comparator.NewComparator(comparator.Size)
	result, err := cmp.Compare(a, b)
	if err != nil {
		fmt.Println("something wrong")
		return
	}

	if result {
		fmt.Printf("Size of book: \"%s\" bigger then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
	} else {
		fmt.Printf("Size of book: \"%s\" smaller then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
	}
}

func exampleComparatorRate(a book.Book, b book.Book) {
	cmp := comparator.NewComparator(comparator.Rate)
	result, err := cmp.Compare(a, b)
	if err != nil {
		fmt.Println("something wrong")
		return
	}

	if result {
		fmt.Printf("Rate of book: \"%s\" bigger then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
	} else {
		fmt.Printf("Rate of book: \"%s\" smaller then book: \"%s\" \n", a.GetTitle(), b.GetTitle())
	}
}
