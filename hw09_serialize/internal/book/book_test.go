package book

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanBookUnmarshal(t *testing.T) {
	expected := *NewBook(8881, "Title1", "Author1", 111, 111, 1.3)

	jsonData := `{
			"id": 8881,
			"title": "Title1",
			"author": "Author1",
			"year": 111,
			"size": 111,
			"rate": 1.3
		}`

	var actual Book
	json.Unmarshal([]byte(jsonData), &actual)

	assert.Equal(t, expected, actual)
}

func TestFailBookUnmarshalWhenJsonWrong(t *testing.T) {
	jsonData := `{
			"id 8881
			"title": "Title1",
			"author": "Author1",
			"year": 111,
			"size": 111,
			"rate": 1.3
		}`

	var actual Book
	err := json.Unmarshal([]byte(jsonData), &actual)

	assert.IsType(t, &json.SyntaxError{}, err)
}

func TestCanBookMarshal(t *testing.T) {
	book := NewBook(8881, "Title1", "Author1", 111, 111, 1.3)

	expected := `{"id": 8881,"title":"Title1","author":"Author1","year": 111,"size":111,"rate": 1.3}`

	bookMarshaled, _ := json.Marshal(book)

	assert.JSONEq(t, expected, string(bookMarshaled))
}

func TestFailBookMarshalWhenDereferenceOperator(t *testing.T) {
	book := *NewBook(8881, "Title1", "Author1", 111, 111, 1.3)

	expected := `{"id": 8881,"title":"Title1","author":"Author1","year": 111,"size":111,"rate": 1.3}`

	//nolint:all
	bookMarshaled, _ := json.Marshal(book)

	assert.NotEqual(t, expected, string(bookMarshaled))
}

func TestCanBooksUnmarshal(t *testing.T) {
	book1 := NewBook(8881, "Title1", "Author1", 111, 111, 1.3)
	book2 := NewBook(8882, "Title2", "Author2", 222, 222, 2.3)
	expectedBooks := NewBooks(book1)
	expectedBooks.Add(book2)

	jsonData := `[
		{
			"id": 8881,
			"title": "Title1",
			"author": "Author1",
			"year": 111,
			"size": 111,
			"rate": 1.3
		},
		{
			"id": 8882,
			"title": "Title2",
			"author": "Author2",
			"year": 222,
			"size": 222,
			"rate": 2.3
		}
	]`

	var actual Books
	json.Unmarshal([]byte(jsonData), &actual)

	assert.Equal(t, *expectedBooks, actual)
}

func TestFailBooksUnmarshalWhenJsonWrong(t *testing.T) {
	book1 := NewBook(8881, "Title1", "Author1", 111, 111, 1.3)
	book2 := NewBook(8882, "Title2", "Author2", 222, 222, 2.3)
	expectedBooks := NewBooks(book1)
	expectedBooks.Add(book2)

	jsonData := `[
		{
			"id 8881,
			"title": "Title1",
			"author": "Author1",
			"year": 111,
			"size": 111,
			"rate": 1.3
		},
		{
			"id": 8882,
			"title": "Title2",
			"auth
			"year": 222,
			"size": 222,
			"rate": 2.3
		}
	]`

	var actual Books
	err := json.Unmarshal([]byte(jsonData), &actual)

	assert.IsType(t, &json.SyntaxError{}, err)
}

func TestCanBooksMarshal(t *testing.T) {
	book1 := NewBook(8881, "Title1", "Author1", 111, 111, 1.3)
	book2 := NewBook(8882, "Title2", "Author2", 222, 222, 2.3)
	books := NewBooks(book1)
	books.Add(book2)

	expected := `[{"id":8881,"title":"Title1","author":"Author1","year":111,"size":111,"rate":1.300000},
	{"id":8882,"title":"Title2","author":"Author2","year":222,"size":222,"rate":2.300000}]`
	expected = strings.ReplaceAll(expected, "\n", "")
	expected = strings.ReplaceAll(expected, "\t", "")

	actual, _ := json.Marshal(books)

	assert.Equal(t, []byte(expected), actual)
}
