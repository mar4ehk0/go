package post

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Service struct {
	repo InMemoryPostRepository
}

func NewPostService(repo InMemoryPostRepository) *Service {
	return &Service{repo: repo}
}

func (p *Service) Create(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
	}

	err = p.repo.persist(post)
	if errors.Is(err, ErrAlreadyExist) {
		w.WriteHeader(http.StatusConflict)
		os.Stdout.Write([]byte("Post already exist.\n"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	os.Stdout.Write([]byte("Post created.\n"))
}

func (p *Service) Read(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")

	post, err := p.repo.getByTitle(title)
	if errors.Is(err, ErrPostNotFound) {
		w.WriteHeader(http.StatusNotFound)
		os.Stdout.Write([]byte("Post not found.\n"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		os.Stdout.Write([]byte("Can not marshal post.\n"))
		return
	}

	os.Stdout.Write([]byte("Post found.\n"))
}
