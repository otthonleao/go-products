package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/otthonleao/go-products.git/internal/dto"
	"github.com/otthonleao/go-products.git/internal/entity"
	"github.com/otthonleao/go-products.git/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (handler *UserHandler) Create(response http.ResponseWriter, request *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	userRequest, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserDB.Create(userRequest)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusCreated)
}