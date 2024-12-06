package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ma4ehk0/go/hw13_http/server/internal/post"
	"github.com/ma4ehk0/go/hw13_http/server/pkg/server"
)

func main() {
	addr := server.NewAddr()

	repo := post.NewInMemoryPostRepository()

	handler := post.NewPostService(repo)

	router := initializeRoutes(handler)

	server := &http.Server{
		Addr:              addr.Connection(),
		Handler:           router,
		ReadHeaderTimeout: time.Second,
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initializeRoutes(p *post.Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts/{title}", p.Read)
	mux.HandleFunc("POST /posts", p.Create)
	return mux
}
