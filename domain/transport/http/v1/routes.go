package v1

import (
	"fmt"
	"net/http"

	"github.com/92Sam/ms-users/domain/controllers"
	"github.com/92Sam/ms-users/domain/transport/http/v1/middlewares"
	"github.com/gorilla/mux"
)

type Routes struct {
	RouterV1 *mux.Router
	*controllers.Controllers
}

func NewRoutes(router *mux.Router, ctrl *controllers.Controllers) {
	a := new(Routes)
	a.Controllers = ctrl
	a.initRoutesV1(router)
}

func (a *Routes) initRoutesV1(router *mux.Router) {
	a.RouterV1 = router.NewRoute().Methods(
		http.MethodPost,
		http.MethodGet,
		http.MethodDelete,
		http.MethodPut,
		http.MethodPatch).PathPrefix("/v1").Subrouter()
	a.initializeRoutesProducts()
	a.initializeRoutesUsers()
	a.initializeRoutesAuth()
}

func (a *Routes) initializeRoutesProducts() {

	fmt.Println("Init Routes Products")
	a.RouterV1.Path("/productsfree").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetProductFree)

	u := a.RouterV1.PathPrefix("/products").Subrouter()
	u.Use(middlewares.MiddlewareUser)
	u.Path("").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetProduct)
	u.Path("").Methods(http.MethodPost).HandlerFunc(a.Controllers.CreateProduct)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetProductById)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodPatch).HandlerFunc(a.Controllers.UpdateProductById)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodDelete).HandlerFunc(a.Controllers.GetProductById)
}

func (a *Routes) initializeRoutesUsers() {
	fmt.Println("Init Routes Users")
	u := a.RouterV1.PathPrefix("/users").Subrouter()

	// u.Use(middleware.Middleware)
	u.Path("").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetUsers)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodDelete).HandlerFunc(a.Controllers.DeleteUserById)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodPatch).HandlerFunc(a.Controllers.UpdateUserById)
	u.Path("/{id:[a-z0-9-]+}").Methods(http.MethodGet).HandlerFunc(a.Controllers.GetUsersById)
}

func (a *Routes) initializeRoutesAuth() {
	fmt.Println("Init Routes Auth")

	u := a.RouterV1.PathPrefix("/auth").Subrouter()
	u.Path("/login").Methods(http.MethodPost).HandlerFunc(a.Controllers.Login)
	u.Path("/signup").Methods(http.MethodPost).HandlerFunc(a.Controllers.Signup)
}
