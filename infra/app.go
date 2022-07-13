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
}

func InitApp(errChanServer chan error, statusChanServer chan bool) {
	a := new(App)
	a.Router = mux.NewRouter()
	ctx := context.Background()
	dynamoDb := initPersistence(ctx)

	a.initRepositories(dynamoDb)
	a.initServices()
	a.initControllers()
	a.initRoutes()

	// Go Routine Start Server v1
	go v1.StartServer(a.Router, errChanServer, statusChanServer)
}

func (a *App) initRoutes() {
	v1.NewRoutes(a.Router, a.Controllers)
}

func (a *App) initRepositories(dynamoSvc *persistence.DynamoDbContext) {
	repo := new(repositories.Repositories)
	repo.ProductRepository = repositories.NewProductRepository(dynamoSvc)
	repo.UserRepository = repositories.NewUserRepository(dynamoSvc)
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
	ctr.ProductController = controllers.NewProductController(a.Services)
	ctr.UserController = controllers.NewUserController(a.Services)
	a.Controllers = ctr
}

func initPersistence(ctx context.Context) *persistence.DynamoDbContext {
	return persistence.InitDynamoDb(ctx)
}
