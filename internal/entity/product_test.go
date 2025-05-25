package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thenopholo/back_commerce/pkg/entity"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Macbook", 1250.99)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Macbook", product.Name)
	assert.Equal(t, 1250.99, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 1250.99)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}
func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Mackbook", 0.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}
func TestProductWhenPriceIsNotValid(t *testing.T) {
	p, err := NewProduct("Mackbook", -1)

	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestValidate(t *testing.T) {
	p, err := NewProduct("Macbook", 1250.99)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())

}

func TestProduct_ValidateErrIDIsRequired(t *testing.T) {
	product := &Product{
		ID:        entity.ID{},
		Name:      "Test Product",
		Price:     10.0,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	assert.Equal(t, ErrIDIsRequired, err)
}
