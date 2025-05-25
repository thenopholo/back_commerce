package db

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenopholo/back_commerce/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{})
	return db
}

func TestCreateProduct(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	product, err := entity.NewProduct("Macbook Pro", 2999.99)
	assert.Nil(t, err)

	err = productRepo.CreateProduct(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindByID(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	// Criar produto
	product, err := entity.NewProduct("iPhone 15", 999.99)
	assert.Nil(t, err)

	err = productRepo.CreateProduct(product)
	assert.Nil(t, err)

	// Buscar por ID
	foundProduct, err := productRepo.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, foundProduct.ID)
	assert.Equal(t, product.Name, foundProduct.Name)
	assert.Equal(t, product.Price, foundProduct.Price)
}

func TestFindByID_NotFound(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	// Buscar por ID inexistente
	_, err := productRepo.FindByID("nonexistent-id")
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestFindAll(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	products := []*entity.Product{}

	product1, _ := entity.NewProduct("Macbook Pro", 2999.99)
	product2, _ := entity.NewProduct("iPhone 15", 999.99)
	product3, _ := entity.NewProduct("iPad Air", 599.99)

	products = append(products, product1, product2, product3)

	for _, p := range products {
		err := productRepo.CreateProduct(p)
		assert.Nil(t, err)
	}


	foundProducts, err := productRepo.FindAll(0, 0, "asc")
	assert.Nil(t, err)
	assert.Len(t, foundProducts, 3)
}

func TestFindAll_WithPagination(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	// Criar 5 produtos
expectedCount := 65
	for i := 1; i <= expectedCount; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*1000+10)
		assert.Nil(t, err)

		err = productRepo.CreateProduct(product)
		assert.Nil(t, err)
	}

	foundProducts, err := productRepo.FindAll(1, 30, "asc")
	assert.Nil(t, err)
	assert.Len(t, foundProducts, 30)
  assert.Equal(t, "Product 1", foundProducts[0].Name)
  assert.Equal(t, "Product 30", foundProducts[29].Name)

	foundProducts, err = productRepo.FindAll(2, 30, "asc")
	assert.Nil(t, err)
	assert.Len(t, foundProducts, 30)
  assert.Equal(t, "Product 31", foundProducts[0].Name)
  assert.Equal(t, "Product 60", foundProducts[29].Name)

	foundProducts, err = productRepo.FindAll(3, 30, "asc")
	assert.Nil(t, err)
	assert.Len(t, foundProducts, 5)
  assert.Equal(t, "Product 61", foundProducts[0].Name)
  assert.Equal(t, "Product 65", foundProducts[4].Name)

}

func TestFindAll_InvalidSort(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	product, _ := entity.NewProduct("Test Product", 100.0)
	err := productRepo.CreateProduct(product)
	assert.Nil(t, err)

	foundProducts, err := productRepo.FindAll(0, 0, "invalid")
	assert.Nil(t, err)
	assert.Len(t, foundProducts, 1)
}

func TestUpdate(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	product, err := entity.NewProduct("Original Name", 100.0)
	assert.Nil(t, err)

	err = productRepo.CreateProduct(product)
	assert.Nil(t, err)

	product.Name = "Updated Name"
	product.Price = 200.0

	err = productRepo.Update(product)
	assert.Nil(t, err)

	updatedProduct, err := productRepo.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, "Updated Name", updatedProduct.Name)
	assert.Equal(t, 200.0, updatedProduct.Price)
}

func TestUpdate_ProductNotFound(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	product, err := entity.NewProduct("Test Product", 100.0)
	assert.Nil(t, err)

	err = productRepo.Update(product)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestDelete(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	product, err := entity.NewProduct("Product to Delete", 100.0)
	assert.Nil(t, err)

	err = productRepo.CreateProduct(product)
	assert.Nil(t, err)

	err = productRepo.Delete(product.ID.String())
	assert.Nil(t, err)

	_, err = productRepo.FindByID(product.ID.String())
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestDelete_ProductNotFound(t *testing.T) {
	db := setupTestDB()
	productRepo := NewProduct(db)

	err := productRepo.Delete("nonexistent-id")
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
