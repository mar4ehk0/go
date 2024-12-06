package post

import (
	"errors"
	"fmt"
)

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	ErrPostNotFound = errors.New("post not found")
	ErrAlreadyExist = errors.New("already exist")
)

type InMemoryPostRepository []Post

func NewInMemoryPostRepository() InMemoryPostRepository {
	return make(InMemoryPostRepository, 0, 100)
}

func (r *InMemoryPostRepository) persist(post Post) error {
	fmt.Println(*r)
	fmt.Printf("%#v\n", post)
	for _, v := range *r {
		if v.Title == post.Title {

			return ErrAlreadyExist
		}
	}
	*r = append(*r, post)

	fmt.Println(*r)
	return nil
}

func (i *InMemoryPostRepository) getByTitle(title string) (Post, error) {
	for _, post := range *i {
		if post.Title == title {

			return post, nil
		}
	}

	return Post{}, ErrPostNotFound
}
