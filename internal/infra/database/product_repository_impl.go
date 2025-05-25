package database

import (
	"github.com/thenopholo/back_commerce/internal/entity"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (p *ProductRepositoryImpl) CreateProduct(product *entity.Product) error {
	return p.DB.Create(&product).Error
}

func (p *ProductRepositoryImpl) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *ProductRepositoryImpl) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *ProductRepositoryImpl) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *ProductRepositoryImpl) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
