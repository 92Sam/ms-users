package services

import (
	"fmt"
	"time"

	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/domain/repositories"
	"github.com/92Sam/ms-users/domain/serializables"
	"github.com/google/uuid"
)

type IProductService interface {
	// GetProductByEmail(loginRequest *serializables.ProductRequest) (*models.Product, error)
	Create(ProductRequest *serializables.ProductRequest) (*models.Product, error)
	GetProducts() ([]*models.Product, error)
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

func (ps *ProductService) Create(productRequest *serializables.ProductRequest) (*models.Product, error) {
	user := &models.Product{
		Id:          uuid.NewString(),
		Name:        productRequest.Name,
		Description: *productRequest.Description,
		Rating:      productRequest.Rating,
		CreateAt:    time.Now(),
	}

	record, err := ps.Repositories.ProductRepository.Create(user)
	if err != nil {
		fmt.Errorf("Error ->", err)
		return nil, err
	}

	return record, nil
}

func (ps *ProductService) GetProducts() ([]*models.Product, error) {

	records, err := ps.Repositories.ProductRepository.GetProducts()
	if err != nil {
		fmt.Errorf("Error ->", err)
		return nil, err
	}

	return records, nil
}
