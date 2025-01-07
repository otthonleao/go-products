package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/otthonleao/go-products.git/internal/dto"
	"github.com/otthonleao/go-products.git/internal/entity"
	"github.com/otthonleao/go-products.git/internal/infra/database"
	entityPkg "github.com/otthonleao/go-products.git/pkg/entity"
)

type ProductHandler struct {
	productDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		productDB: db,
	}
}

// Create Product godoc
// @Summary     Create a product
// @Description Create a new product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       request     body    dto.CreateProductInput     true    "Product request"
// @Success     201
// @Failure     500	 {object}    Error
// @Router      /products    [post]
// @Security    ApiKeyAuth
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

// Get Product godoc
// @Summary     Get a product
// @Description Get a product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id          path    string     true    "Product ID"		Format(uuid)
// @Success     200		{object}    entity.Product
// @Failure     404
// @Failure     500		{object}    Error
// @Router      /products/{id}    [get]
// @Security    ApiKeyAuth
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

// List Products godoc
// @Summary     List products
// @Description Get all products
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       page        query    int     false    "Page number"
// @Param       limit       query    int     false    "Number of items per page"
// @Success     200		{array}    entity.Product
// @Failure     500		{object}    Error
// @Router      /products    [get]
// @Security    ApiKeyAuth
func (handler *ProductHandler) GetProducts(response http.ResponseWriter, request *http.Request) {
	
	page := request.URL.Query().Get("page")
	limit := request.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := request.URL.Query().Get("sort")

	products, err := handler.productDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(products)
}

// Update Product godoc
// @Summary     Update a product
// @Description Update a product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id          path    string     true    "Product ID"		Format(uuid)
// @Param       request     body    dto.CreateProductInput     true    "Product request"
// @Success     200
// @Failure     404
// @Failure     500		{object}    Error
// @Router      /products/{id}    [put]
// @Security    ApiKeyAuth
func (handler *ProductHandler) UpdateProduct(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	
	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = handler.productDB.FindByID(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	err = handler.productDB.Update(&product)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Delete Product godoc
// @Summary     Delete a product
// @Description Delete a product by ID
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id          path    string     true    "Product ID"		Format(uuid)
// @Success     204
// @Failure     404
// @Failure     500		{object}    Error
// @Router      /products/{id}    [delete]
// @Security    ApiKeyAuth
func (handler *ProductHandler) DeleteProduct(response http.ResponseWriter, request *http.Request) {
	
	id := chi.URLParam(request, "id")
	
	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := handler.productDB.FindByID(id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	err = handler.productDB.Delete(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}