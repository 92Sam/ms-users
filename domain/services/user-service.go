package services

import (
	"fmt"
	"time"

	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/domain/repositories"
	"github.com/92Sam/ms-users/domain/serializables"
	"github.com/google/uuid"
)

type UserService struct {
	*repositories.Repositories
}

func NewUserService(reps *repositories.Repositories) *UserService {
	return &UserService{reps}
}

func (ps *UserService) GetUserByEmail(loginRequest *serializables.AuthUserLoginRequest) (*models.User, error) {
	user, err := ps.Repositories.UserRepository.GetByEmail(loginRequest.Email)
	if err != nil {
		fmt.Errorf("Error ->", err)
		return nil, err
	}

	return user, nil
}

func (ps *UserService) Create(userRequest *serializables.UserRequest) (*models.User, error) {

	user := &models.User{
		Id:       uuid.NewString(),
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.PasswordHashed(),
		CreateAt: time.Now(),
	}

	user, err := ps.Repositories.UserRepository.Create(user)
	if err != nil {
		fmt.Errorf("Error ->", err)
		return nil, err
	}

	return user, nil
}

func (ps *AuthService) GetUsers() ([]*models.User, error) {
	user, err := ps.Repositories.UserRepository.GetUsers()
	if err != nil {
		fmt.Errorf("Error ->", err)
		return nil, err
	}

	return user, nil
}
