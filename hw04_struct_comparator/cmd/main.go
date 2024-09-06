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

	modes := []comparator.ModeEnum{comparator.Year, comparator.Size, comparator.Rate}
	for _, mode := range modes {
		dto := checker.NewDto(
			mode,
			book1,
			book2,
		)
		checker.Check(dto)
	}
}
