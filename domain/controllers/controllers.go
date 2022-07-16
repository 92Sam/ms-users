package controllers

type Controllers struct {
	IProductController
	IAuthUserController
}

type IAuthUserController interface {
	IUserController
	IAuthController
}
