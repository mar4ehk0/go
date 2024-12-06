package post

import (
	"errors"
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

func (i *InMemoryPostRepository) persist(post Post) error {
	for _, v := range *i {
		if v.Title == post.Title {
			return ErrAlreadyExist
		}
	}
	*i = append(*i, post)

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
