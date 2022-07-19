package infra

import (
	"context"

	"github.com/92Sam/ms-users/domain/controllers"
	"github.com/92Sam/ms-users/domain/persistence"
	"github.com/92Sam/ms-users/domain/repositories"
	"github.com/92Sam/ms-users/domain/services"
	v1 "github.com/92Sam/ms-users/domain/transport/http/v1"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	*repositories.Repositories
	*controllers.Controllers
	*services.Services
	*AppPersistence
}

type AppPersistence struct {
	*persistence.DynamoDbContext
	*persistence.Neo4jDbContext
}

func InitApp(errChanServer chan error, statusChanServer chan bool) {
	a := new(App)
	a.Router = mux.NewRouter()
	ctx := context.Background()

	a.initPersistence(ctx)
	a.initRepositories()
	a.initServices()
	a.initControllers()
	a.initRoutes()

	// Go Routine Start Server v1
	go v1.StartServer(a.Router, errChanServer, statusChanServer)
}

func (a *App) initRoutes() {
	v1.NewRoutes(a.Router, a.Controllers)
}

func (a *App) initRepositories() {
	repo := new(repositories.Repositories)
	repo.ProductRepository = repositories.NewProductRepository(a.Neo4jDbContext)
	repo.UserRepository = repositories.NewUserRepository(a.DynamoDbContext)
	a.Repositories = repo
}

func (a *App) initServices() {
	srv := new(services.Services)
	srv.ProductService = services.NewProductService(a.Repositories)
	srv.UserService = services.NewUserService(a.Repositories)
	srv.AuthService = services.NewAuthService(a.Repositories)
	a.Services = srv
}

func (a *App) initControllers() {
	ctr := new(controllers.Controllers)
	ctr.IProductController = controllers.NewProductController(a.Services)
	ctr.IAuthUserController = controllers.NewUserController(a.Services)
	a.Controllers = ctr
}

func (a *App) initPersistence(ctx context.Context) {
	a.AppPersistence = &AppPersistence{
		persistence.InitDynamoDb(ctx),
		persistence.InitNeo4jDb(),
	}
}
