package main

import (
	"log"
	"net/http"

	"github.com/incognito/store-api/internal/categories"
	"github.com/incognito/store-api/internal/config"
	"github.com/incognito/store-api/internal/database"
	"github.com/incognito/store-api/internal/products"
)

func main() {
	cfg := config.LoadConfig()

	//log.Println("DB URL:", cfg.Database_URL)
	//log.Println("DB PORT:", cfg.Port)

	db, err := database.NewPostgresCon(cfg.Database_URL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//products
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	handler := products.NewHandler(serv)

	//categories
	crepo := categories.NewRepository(db)
	cserv := categories.NewService(crepo)
	chandler := categories.NewHandler(cserv)

	mux := http.NewServeMux()
	//products
	mux.HandleFunc("GET /products", handler.GetProductHandler)
	mux.HandleFunc("POST /products", handler.AddProductHandler)
	mux.HandleFunc("GET /products/{id}", handler.GetProductByIdHandler)
	mux.HandleFunc("PUT /products/{id}", handler.UpdateProductHandler)
	mux.HandleFunc("DELETE /products/{id}", handler.DeleteProductHandler)

	//categories
	mux.HandleFunc("POST /categories", chandler.CreateCategoryHandler)
	mux.HandleFunc("GET /categories", chandler.GetCategoryHandler)
	mux.HandleFunc("GET /categories/{id}", chandler.GetCategoryByIdHandler)
	mux.HandleFunc("PUT /categories/{id}", chandler.UpdateCategoryHandler)

	log.Printf("Server running on :%s", cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
