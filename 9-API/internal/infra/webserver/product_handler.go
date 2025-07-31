package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/sancheschris/goexpert/9-APIS/internal/dto"
	"github.com/sancheschris/goexpert/9-APIS/internal/entity"
	"github.com/sancheschris/goexpert/9-APIS/internal/infra/database"
)


type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, "Error saving product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}