package post

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type PostService struct {
	repo InMemoryPostRepository
}

func NewPostService(repo InMemoryPostRepository) *PostService {
	return &PostService{repo: repo}
}

func (p *PostService) Create(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
	}

	err = p.repo.persist(post)
	if err == ErrAlreadyExist {
		w.WriteHeader(http.StatusConflict)
		os.Stdout.Write([]byte("Post already exist.\n"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	os.Stdout.Write([]byte("Post created.\n"))
}

func (p *PostService) Read(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")

	post, err := p.repo.getByTitle(title)
	if err == ErrPostNotFound {
		w.WriteHeader(http.StatusNotFound)
		os.Stdout.Write([]byte("Post not found.\n"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

	os.Stdout.Write([]byte("Post found.\n"))
}
