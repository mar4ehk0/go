package main

import (
	"log"
	"net/http"

	"github.com/ma4ehk0/go/hw13_http/server/internal/post"
	"github.com/ma4ehk0/go/hw13_http/server/pkg/server"
)

func main() {
	addr := server.NewAddr()

	repo := post.NewInMemoryPostRepository()

	handler := post.NewPostService(repo)

	router := initializeRoutes(handler)

	server := &http.Server{
		Addr:    addr.Connection(),
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

func initializeRoutes(p *post.PostService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts/{title}", p.Read)
	mux.HandleFunc("POST /posts", p.Create)
	return mux
}
