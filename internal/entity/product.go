package entity

import (
	"errors"
	"time"

	"github.com/thenopholo/back_commerce/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("ID is required")
	ErrInvalidID       = errors.New("invalid ID")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID == (entity.ID{}) {
		return ErrIDIsRequired
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0.0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0.0 {
		return ErrInvalidPrice
	}

	return nil
}
