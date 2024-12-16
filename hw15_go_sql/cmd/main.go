package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/product"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/handler"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	godotenv.Load()

	addr := server.NewAddr()

	db, err := sqlx.Connect("pgx", os.Getenv("APP_DB_DSN"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	repoProduct := product.NewRepo(db)
	serviceProduct := product.NewService(repoProduct)
	handlerProduct := handler.NewHandler(serviceProduct)

	router := initializeRoutes(handlerProduct)

	server := &http.Server{
		Addr:              addr.Connection(),
		Handler:           router,
		ReadHeaderTimeout: time.Second,
	}
	log.Println("Listening...")
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func initializeRoutes(h *handler.Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("GET /products/{id}", h.GetProductById)
	mux.HandleFunc("PATCH /products/{id}", h.UpdateProductById)
	return mux
}
