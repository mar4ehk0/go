package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/product"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/user"
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

	mux := http.NewServeMux()

	repoProduct := product.NewRepo(db)
	serviceProduct := product.NewService(repoProduct)
	handlerProduct := product.NewHandler(serviceProduct)
	handlerProduct.InitializeRoutes(mux)

	repoUser := user.NewRepo(db)
	serviceUser := user.NewService(repoUser)
	handlerUser := user.NewHandler(serviceUser)
	handlerUser.InitializeRoutes(mux)


	server := &http.Server{
		Addr:              addr.Connection(),
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}
	log.Println("Listening...")
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// func initializeRoutes(p *product.Handler) http.Handler {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("POST /products", p.CreateProduct)
// 	mux.HandleFunc("GET /products/{id}", p.GetProductById)
// 	mux.HandleFunc("PATCH /products/{id}", p.UpdateProductById)
// 	return mux
// }
