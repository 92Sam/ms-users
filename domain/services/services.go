package services

type Services struct {
	UserService    IUserService
	AuthService    IAuthService
	ProductService IProductService
}
