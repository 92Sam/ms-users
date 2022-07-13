package controllers

import (
	"fmt"
	"net/http"

	"github.com/92Sam/ms-users/domain/services"
	"github.com/gorilla/mux"
)

type ProductController struct {
	*services.Services
}

func NewProductController(svc *services.Services) *ProductController {
	return &ProductController{svc}
}

func (a *ProductController) GetProductFree(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars)
	fmt.Println("Aqui GetProductFree")

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
	vars := mux.Vars(r)
	fmt.Print(vars)
	fmt.Println("Aqui GetProduct")
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