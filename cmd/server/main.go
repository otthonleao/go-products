package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	"github.com/otthonleao/go-products.git/configs"
	_ "github.com/otthonleao/go-products.git/docs"
	"github.com/otthonleao/go-products.git/internal/entity"
	"github.com/otthonleao/go-products.git/internal/infra/database"
	"github.com/otthonleao/go-products.git/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title			Go Products API
// @version			1.0
// @description		Product API with authentication in GO
// @termsOfService	http://swagger.io/terms/

// @contact.name	Otthon Leão
// @contact.url		http://meusite.com
// @contact.email	otthonleao@hotmail.com

// @license.name	MIT
// @license.url		http://mit.com

// @host						localhost:8000
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	// Carregar configurações
	configs, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	// Inicializar roteador
	route := chi.NewRouter()
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)
	route.Use(middleware.WithValue("jwt", configs.TokenAuth))
	route.Use((middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn)))

	// Register the handler
	route.Route("/products", func(chiRoute chi.Router) {
		chiRoute.Use(jwtauth.Verifier(configs.TokenAuth))
		chiRoute.Use(jwtauth.Authenticator)
		chiRoute.Post("/", productHandler.Create)
		chiRoute.Get("/{id}", productHandler.GetProduct)
		chiRoute.Get("/", productHandler.GetProducts)
		chiRoute.Put("/{id}", productHandler.UpdateProduct)
		chiRoute.Delete("/{id}", productHandler.DeleteProduct)
	})

	route.Post("/users", userHandler.Create)
	route.Post("/users/login", userHandler.GetJWT)

	// Subindo a documentação do webservice
	route.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/swagger/doc.json")))


	// http.HandleFunc("/products", productHandler.Create)
	http.ListenAndServe(":8000", route)
}
