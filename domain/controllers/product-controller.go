package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/92Sam/ms-users/domain/serializables"
	"github.com/92Sam/ms-users/domain/services"
	"github.com/92Sam/ms-users/utils"
	"github.com/gorilla/mux"
)

type IProductController interface {
	GetProductFree(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	UpdateProductById(w http.ResponseWriter, r *http.Request)
}

type ProductController struct {
	*services.Services
}

func NewProductController(svc *services.Services) IProductController {
	return &ProductController{svc}
}

func (a *ProductController) GetProductFree(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars)
	fmt.Println("GetProductFree")

	// id, err := strconv.Atoi(vars["id"])
	// if err != nil {
	//     respondWithError(w, http.StatusBadRequest, "Invalid product ID")
	//     return
	// }

	// p := product{ID: id}
	// if err := p.getProduct(a.DB); err != nil {
	//     switch err {
	//     case sql.ErrNoRows:
	//         respondWithError(w, http.StatusNotFound, "Product not found")
	//     default:
	//         respondWithError(w, http.StatusInternalServerError, err.Error())
	//     }
	//     return
	// }

	// respondWithJSON(w, http.StatusOK, p)
}

func (a *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {

	produtList, err := a.Services.ProductService.GetProducts()
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, produtList)
	return
}

func (a *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := &serializables.ProductRequest{}
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := a.Services.ProductService.Create(p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// dtoResponseList := make([]*serializables.UserResponse, len(resp))
	// for k, v := range resp {
	// 	dtoResponseList[k] = &serializables.UserResponse{
	// 		Id:        v.Id,
	// 		Name:      v.Name,
	// 		Email:     v.Email,
	// 		CreatedAt: v.CreateAt,
	// 	}
	// }

	utils.RespondWithJSON(w, http.StatusOK, product)
	return
}

func (a *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars)
	fmt.Println("GetProductById")
}

func (a *ProductController) UpdateProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars)
	fmt.Println("UpdateProductById")
}
