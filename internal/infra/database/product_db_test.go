package database

import (
	"fmt"
	"math/rand"
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

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product Test %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	// Testando página 1
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product Test 1", products[0].Name)
	assert.Equal(t, "Product Test 10", products[9].Name)
	// Testando página 2
	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product Test 11", products[0].Name)
	assert.Equal(t, "Product Test 20", products[9].Name)
	// Testando página 3 que só pode ter 3 registros de produtos
	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product Test 21", products[0].Name)
	assert.Equal(t, "Product Test 23", products[2].Name)
}


func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product Test", 10)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product Test by ID", 10.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	product.Name = "Product Test Updated"
	product.Price = 20.00
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}