package post

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanInMemoryPostRepositoryPersist(t *testing.T) {
	post := Post{Title: "lorem", Body: "ipsum"}
	post1 := Post{Title: "lorem1", Body: "ipsum1"}
	post2 := Post{Title: "lorem2", Body: "ipsum2"}
	repo := NewInMemoryPostRepository()

	err := repo.persist(post)
	assert.Nil(t, err)
	err = repo.persist(post1)
	assert.Nil(t, err)
	err = repo.persist(post2)
	assert.Nil(t, err)

	assert.Equal(t, repo[0].Title, post.Title)
	assert.Equal(t, repo[0].Body, post.Body)
	assert.Equal(t, repo[1].Title, post1.Title)
	assert.Equal(t, repo[1].Body, post1.Body)
	assert.Equal(t, repo[2].Title, post2.Title)
	assert.Equal(t, repo[2].Body, post2.Body)
}

func TestFailInMemoryPostRepositoryPersist(t *testing.T) {
	post := Post{Title: "lorem", Body: "ipsum"}
	post1 := Post{Title: "lorem", Body: "ipsum1"}
	repo := NewInMemoryPostRepository()

	err := repo.persist(post)
	assert.Nil(t, err)
	err = repo.persist(post1)
	assert.ErrorIs(t, ErrAlreadyExist, err)
}

func TestCanInMemoryPostRepositoryGetByTitle(t *testing.T) {
	post := Post{Title: "lorem", Body: "ipsum"}
	post1 := Post{Title: "lorem1", Body: "ipsum1"}
	post2 := Post{Title: "lorem2", Body: "ipsum2"}
	repo := NewInMemoryPostRepository()
	err := repo.persist(post)
	assert.Nil(t, err)
	err = repo.persist(post1)
	assert.Nil(t, err)
	err = repo.persist(post2)
	assert.Nil(t, err)

	postFromRepo, err := repo.getByTitle("lorem1")
	assert.Nil(t, err)

	assert.Equal(t, post1.Title, postFromRepo.Title)
	assert.Equal(t, post1.Body, postFromRepo.Body)
}

func TestFailInMemoryPostRepositoryGetByTitle(t *testing.T) {
	post := Post{Title: "lorem", Body: "ipsum"}
	post1 := Post{Title: "lorem1", Body: "ipsum1"}
	post2 := Post{Title: "lorem2", Body: "ipsum2"}
	repo := NewInMemoryPostRepository()
	err := repo.persist(post)
	assert.Nil(t, err)
	err = repo.persist(post1)
	assert.Nil(t, err)
	err = repo.persist(post2)
	assert.Nil(t, err)

	postFromRepo, err := repo.getByTitle("notfound")

	assert.Equal(t, "", postFromRepo.Title)
	assert.Equal(t, "", postFromRepo.Body)
	assert.ErrorIs(t, ErrPostNotFound, err)
}
