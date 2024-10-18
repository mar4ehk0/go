package main

import (
	"encoding/json"
	"fmt"

	"github.com/mar4ehk0/hw09_serialize/internal/book"
	pb "github.com/mar4ehk0/hw09_serialize/pkg/book/api/book"
	"google.golang.org/protobuf/proto"
)

func main() {
	var err error

	book1 := book.NewBook(8881, "Title1", "Author1", 111, 111, 1.3)
	book2 := book.NewBook(8882, "Title2", "Author2", 222, 222, 2.3)

	// книга
	book1Marshaled, _ := json.Marshal(book1)
	fmt.Printf("%#v\n", book1Marshaled)

	var bookUnmarshaled book.Book
	err = json.Unmarshal(book1Marshaled, &bookUnmarshaled)
	fmt.Printf("%#v\n", err)
	fmt.Printf("%#v\n", bookUnmarshaled)

	// книги
	books := book.NewBooks(book1)
	books.Add(book2)

	booksMarshaled, _ := json.Marshal(books)

	var booksUnmarshaled book.Books
	err = json.Unmarshal(booksMarshaled, &booksUnmarshaled)
	fmt.Printf("%#v\n", err)
	fmt.Printf("%#v\n", booksUnmarshaled)

	// protobuf
	booksPb := &pb.Books{
		Items: []*pb.Book{
			{
				Id:     8881,
				Title:  "Title1",
				Author: "Author1",
				Year:   111,
				Size:   111,
				Rate:   1.3,
			},
			{
				Id:     8882,
				Title:  "Title2",
				Author: "Author2",
				Year:   222,
				Size:   222,
				Rate:   2.3,
			},
		},
	}

	data, err := proto.Marshal(booksPb)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	newBooks := &pb.Books{}
	err = proto.Unmarshal(data, newBooks)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Println(newBooks)
}
