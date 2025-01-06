package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/otthonleao/go-products.git/internal/dto"
	"github.com/otthonleao/go-products.git/internal/entity"
	"github.com/otthonleao/go-products.git/internal/infra/database"
)

type ProductHandler struct {
	productDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		productDB: db,
	}
}

func (handler *ProductHandler) Create(response http.ResponseWriter, request *http.Request) {
	
	var product dto.CreateProductInput

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	} 

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.productDB.Create(p)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *ProductHandler) GetProduct(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	
	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := handler.productDB.FindByID(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(product)
}