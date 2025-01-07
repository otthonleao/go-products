package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/otthonleao/go-products.git/internal/dto"
	"github.com/otthonleao/go-products.git/internal/entity"
	"github.com/otthonleao/go-products.git/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
	}
}


func (handler *UserHandler) GetJWT(response http.ResponseWriter, request *http.Request) {
	
	jwt := request.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := request.Context().Value("jwtExpiresIn").(int)

	var user dto.GetJWTInput

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	userRequest, err := handler.UserDB.FindByEmail(user.Email)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if !userRequest.CheckPassword(user.Password) {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": userRequest.ID.String(),
		"exp": time.Now().Add(time.Hour * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(accessToken)
}

// Create user godoc
// @Summary		Create a new user
// @Description	Create a new user
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		request		body	dto.CreateUserInput	true	"User request"
// @Success		201
// @Failure		500		{object}	Error
// @Router		/users	[post]
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
		error := Error{Message: err.Error()}
		json.NewEncoder(response).Encode(error)
		return
	}

	err = handler.UserDB.Create(userRequest)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(response).Encode(error)
		return
	}
	response.WriteHeader(http.StatusCreated)
}
