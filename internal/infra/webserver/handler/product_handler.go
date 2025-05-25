package handler

import (
	"encoding/json"
	"net/http"

	"github.com/thenopholo/back_commerce/internal/dto"
	"github.com/thenopholo/back_commerce/internal/entity"
	"github.com/thenopholo/back_commerce/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductRepository
}

func NewProductHandler(db database.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.CreateProduct(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
