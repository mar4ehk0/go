package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/order"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/product"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/user"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		exit()
	}

	addr := server.NewAddr()

	sqlxDB, err := sqlx.Connect("pgx", os.Getenv("APP_DB_DSN"))
	if err != nil {
		log.Println(err)
		exit()
	}
	defer sqlxDB.Close()

	dbConnect := db.NewDBConnect(sqlxDB)

	mux := http.NewServeMux()

	repoProduct := product.NewRepoProduct(dbConnect)
	serviceProduct := product.NewService(repoProduct)
	handlerProduct := product.NewHandler(serviceProduct)
	handlerProduct.InitializeRoutes(mux)

	repoUser := user.NewRepoUser(dbConnect)
	serviceUser := user.NewService(repoUser)
	handlerUser := user.NewHandler(serviceUser)
	handlerUser.InitializeRoutes(mux)

	repoOrder := order.NewRepoOrder(dbConnect)
	serviceOrder := order.NewService(repoOrder, repoProduct, repoUser)
	handlerOrder := order.NewHandler(serviceOrder)
	handlerOrder.InitializeRoutes(mux)

	server := &http.Server{
		Addr:              addr.Connection(),
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}
	log.Println("Listening...")
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		exit()
	}
}

func exit() {
	os.Exit(1)
}
