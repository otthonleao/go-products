package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/otthonleao/go-products.git/configs"
	"github.com/otthonleao/go-products.git/internal/entity"
	"github.com/otthonleao/go-products.git/internal/infra/database"
	"github.com/otthonleao/go-products.git/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Carregar configurações
	_, err := configs.LoadConfig(".");
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Inicializar banco de dados SQLite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	// Inicializar roteador
	route := chi.NewRouter()
	route.Use(middleware.Logger)
	// Register the handler
	route.Post("/products", productHandler.Create)
	route.Get("/products/{id}", productHandler.GetProduct)
	route.Get("/products", productHandler.GetProducts)
	route.Put("/products/{id}", productHandler.UpdateProduct)
	route.Delete("/products/{id}", productHandler.DeleteProduct)

	http.HandleFunc("/products", productHandler.Create)
	http.ListenAndServe(":8000", route)
}