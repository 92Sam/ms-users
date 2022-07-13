package services

import (
	"github.com/92Sam/ms-users/domain/repositories"
)

type CrudImpl interface {
	create(request interface{}) interface{}
	getAll(request interface{}) []interface{}
	update(request interface{}) interface{}
	delete(request interface{}) interface{}
	getById(request interface{}) interface{}
}

type ProductService struct {
	*repositories.Repositories
}

func NewProductService(reps *repositories.Repositories) *ProductService {
	return &ProductService{reps}
}

func (ps *ProductService) create(request interface{}) interface{} {

	return ""
}
