package post

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
	ErrAlreadyExist = errors.New("already exist")
)

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) persist(post Post) error {

	return nil
}

func (i *Repo) getByTitle(title string) (Post, error) {

	return Post{}, ErrPostNotFound
}
