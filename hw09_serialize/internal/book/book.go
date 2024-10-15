package book

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Book struct {
	id     uint
	title  string
	author string
	year   uint
	size   uint
	rate   float64
}

func NewBook(id uint, title string, author string, year uint, size uint, rate float64) *Book {
	book := &Book{id: id, title: title, author: author, year: year, size: size, rate: rate}

	return book
}

func (b *Book) MarshalJSON() ([]byte, error) {
	bookMarshaled := []byte(
		`{"id":` + strconv.Itoa(int(b.id)) + `,
		 "title":"` + b.title + `",
		 "author":"` + b.author + `",
		 "year": ` + strconv.Itoa(int(b.year)) + `,
		 "size": ` + strconv.Itoa(int(b.size)) + `,
		 "rate": ` + strconv.FormatFloat(float64(b.rate), 'f', 6, 64) +
			`}`)

	return bookMarshaled, nil
}

func (b *Book) UnmarshalJSON(value []byte) error {

	var buf map[string]interface{}

	err := json.Unmarshal(value, &buf)
	if err != nil {
		return err
	}

	id, ok := buf["id"]
	if !ok {
		return errors.New("not id")
	}
	b.id, err = convertUint(id)
	if err != nil {
		return err
	}

	author, ok := buf["author"]
	if !ok {
		return errors.New("not author")
	}
	b.author = author.(string)

	title, ok := buf["title"]
	if !ok {
		return errors.New("not title")
	}
	b.title = title.(string)

	year, ok := buf["year"]
	if !ok {
		return errors.New("not year")
	}
	b.year, err = convertUint(year)
	if err != nil {
		return err
	}

	size, ok := buf["size"]
	if !ok {
		return errors.New("not size")
	}
	b.size, err = convertUint(size)
	if err != nil {
		return err
	}

	rate, ok := buf["rate"]
	if !ok {
		return errors.New("not rate")
	}
	b.rate = rate.(float64)

	return nil
}

func convertUint(value any) (uint, error) {
	float64, ok := value.(float64)

	if !ok {
		return 0, errors.New("not convert to uint")
	}

	result := uint(float64)

	return result, nil
}

// func (b *Book) Reset() {
// 	*b = Book{}
// }

// func (b *Book) String() string {
// 	return fmt.Sprintf("Book{id: %d, title: %s, author: %s, year: %d, size: %d, rate: %f}", b.id, b.title, b.author, b.year, b.size, b.rate)
// }

// func (b *Book) ProtoMessage() {}
