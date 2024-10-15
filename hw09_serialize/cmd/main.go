package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/mar4ehk0/hw09_serialize/internal/book"
	pb "github.com/mar4ehk0/hw09_serialize/pkg/book/api/book"
)

func main() {
	var bookUnmarshaled book.Book

	book1 := book.NewBook(8881, "Title", "Author", 2024, 111, 5.3)

	book1Marshaled, errMarshal := json.Marshal(book1)
	// fmt.Printf("%#v\n", book1Marshaled)
	fmt.Printf("%#v\n", errMarshal)

	errUnmarshal := json.Unmarshal(book1Marshaled, &bookUnmarshaled)
	fmt.Printf("%#v\n", errUnmarshal)
	fmt.Printf("%#v\n", bookUnmarshaled)

	bookPb := &pb.Book{
		Id:     8882,
		Title:  "Title2",
		Author: "Author2",
		Year:   2024,
		Size:   222,
		Rate:   4.5,
	}

	data, err := proto.Marshal(bookPb)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
