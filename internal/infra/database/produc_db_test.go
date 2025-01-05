package database

import (
	"testing"

	"github.com/otthonleao/go-products.git/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
		t.Fatalf("Um erro '%s' não era esperado ao abrir uma conexão stub com o banco de dados.", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product Test", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

}