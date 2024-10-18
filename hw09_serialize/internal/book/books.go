package book

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Books struct {
	items []*Book
}

func NewBooks(book *Book) *Books {
	books := Books{}
	books.items = append(books.items, book)

	return &books
}

func (b *Books) Add(book *Book) {
	b.items = append(b.items, book)
}

func (b *Books) MarshalJSON() ([]byte, error) {
	booksMarshaled := make([]string, 0)

	result := `[`

	for _, v := range b.items {
		item := v
		bookMarshaled, _ := json.Marshal(item)
		booksMarshaled = append(booksMarshaled, string(bookMarshaled))
	}

	result += strings.Join(booksMarshaled, ", ") + `]`

	return []byte(result), nil
}

func (b *Books) UnmarshalJSON(value []byte) error {
	// data := []rune(string(value))
	// var buf []rune
	// var bookUnmarshaled *Book

	// for _, v := range data {
	// 	if v == '[' || v == ']' {
	// 		continue
	// 	}
	// 	if v == '{' {
	// 		buf = nil
	// 	}

	// 	buf = append(buf, v)

	// 	if v == '}' {
	// 		item := []byte(string(buf))
	// 		err := json.Unmarshal(item, &bookUnmarshaled)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		b.items = append(b.items, bookUnmarshaled)
	// 		bookUnmarshaled = nil
	// 	}
	// }

	var data []map[string]interface{}

	errData := json.Unmarshal(value, &data)
	if errData != nil {
		return errData
	}
	var bookUnmarshaled *Book

	for _, v := range data {
		itemSrc, _ := json.Marshal(v)
		fmt.Println(itemSrc)

		errItemSrc := json.Unmarshal(itemSrc, &bookUnmarshaled)
		if errItemSrc != nil {
			return errItemSrc
		}
		b.items = append(b.items, bookUnmarshaled)
		bookUnmarshaled = nil
	}

	return nil
}
