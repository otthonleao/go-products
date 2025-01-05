package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Otthon Leão", "test@mail.com", "senha123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Otthon Leão", user.Name)
	assert.Equal(t, "test@mail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Otthon Leão", "test@mail.com", "senha123")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("senha123"))
	assert.False(t, user.CheckPassword("senha1234"))
}