package database

import "github.com/otthonleao/go-products.git/internal/entity"

type UserInterface interface {
	Create(user *entity.User)
	FindByEmail(email string) (*entity.User, error)
}