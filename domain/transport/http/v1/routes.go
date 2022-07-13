package v1

import (
	"fmt"
	"net/http"

	"github.com/92Sam/ms-users/domain/controllers"
	"github.com/gorilla/mux"
)

type Routes struct {
	Router *mux.Router
	*controllers.Controllers
}

func NewRoutes(router *mux.Router, ctrl *controllers.Controllers) {
	a := new(Routes)
	a.Router = router
	a.Controllers = ctrl

	a.initializeRoutesProducts()
	a.initializeRoutesUsers()
	a.initializeRoutesAuth()
}

func (a *Routes) initializeRoutesProducts() {

	fmt.Println("Init Routes Products")
	a.Router.Path("/productsfree").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetProductFree)

	u := a.Router.PathPrefix("/products-commission").Subrouter()
	// u.Use(middlewares)
	u.Path("/products").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetProduct)

	// a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	// a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *Routes) initializeRoutesUsers() {
	fmt.Println("Init Routes Users")
	u := a.Router.PathPrefix("/users").Subrouter()

	// u.Use(middleware.Middleware)
	u.Path("").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetUsers)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodDelete).HandlerFunc(a.Controllers.DeleteUserById)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodPatch).HandlerFunc(a.Controllers.UpdateUserById)
}

func (a *Routes) initializeRoutesAuth() {
	fmt.Println("Init Routes Auth")

	u := a.Router.PathPrefix("/auth").Subrouter()
	u.Path("/login").Methods(http.MethodPost).HandlerFunc(a.Controllers.Login)
	u.Path("/signup").Methods(http.MethodPost).HandlerFunc(a.Controllers.Signup)
}
