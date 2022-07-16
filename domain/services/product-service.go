package services

import (
	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/domain/repositories"
	"github.com/92Sam/ms-users/domain/serializables"
)

type IProductService interface {
	// GetProductByEmail(loginRequest *serializables.ProductRequest) (*models.Product, error)
	Create(ProductRequest *serializables.ProductRequest) (*models.Product, error)
	// GetProducts() ([]*models.Product, error)
	// GetProductsById(id string) (*models.Product, error)
	// DeleteProductById(id string) (bool, error)
	// UpdateProductById(id string, ProductReq *serializables.UserRequest) (bool, error)
}

type ProductService struct {
	*repositories.Repositories
}

func NewProductService(reps *repositories.Repositories) IProductService {
	return &ProductService{reps}
}

func (ps *ProductService) Create(ProductRequest *serializables.ProductRequest) (*models.Product, error) {
	return new(models.Product), nil
}
