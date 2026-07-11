package main

import (
	"log"
	"net/http"

	"github.com/incognito/store-api/internal/config"
	"github.com/incognito/store-api/internal/database"
	"github.com/incognito/store-api/internal/products"
)

func main() {
	cfg := config.LoadConfig()

	log.Println("DB URL:", cfg.Database_URL)
	log.Println("DB PORT:", cfg.Port)

	db, err := database.NewPostgresCon(cfg.Database_URL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	handler := products.NewHandler(serv)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", handler.GetProductHandler)
	mux.HandleFunc("POST /products", handler.AddProductHandler)
	mux.HandleFunc("GET /products/{id}", handler.GetProductByIdHandler)

	log.Printf("Server running on :%s", cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
