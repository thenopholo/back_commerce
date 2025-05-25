package db

import "github.com/thenopholo/back_commerce/internal/entity"

type ProductRepository interface {
  CreateProduct(product *entity.Product) error
  FindAll(page, limit int, sort string) ([]entity.Product, error)
  FindByID(id string) (*entity.Product, error)
  Update(product *entity.Product) error
  Delete(id string) error
}