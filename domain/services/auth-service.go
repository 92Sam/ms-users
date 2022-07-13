package services

import (
	"fmt"
	"time"

	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/domain/repositories"
	"github.com/92Sam/ms-users/domain/serializables"
	"github.com/google/uuid"
)

type AuthService struct {
	*repositories.Repositories
}

func NewAuthService(reps *repositories.Repositories) *AuthService {
	return &AuthService{reps}
}

func (ps *AuthService) Login(request *serializables.AuthUserLoginRequest) (*models.User, error) {

	user, err := ps.Repositories.UserRepository.GetByEmail(request.Email)
	if err != nil {
		fmt.Errorf("Error ->", err)
		return nil, err
	}

	err = user.CompareHashAndPassword(request.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ps *AuthService) Signup(request *serializables.AuthUserSignupRequest) (*models.User, error) {

	err := ps.isUserMailExist(request.Email)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Id:       uuid.NewString(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.PasswordHashed(),
		CreateAt: time.Now(),
	}

	user, err = ps.Repositories.UserRepository.Create(user)
	if err != nil {
		fmt.Errorf("Error ->", err)
		return nil, err
	}

	return user, nil
}

func (ps *AuthService) isUserMailExist(email string) error {

	user, err := ps.Repositories.UserRepository.GetByEmail(email)
	if err != nil {
		fmt.Errorf("Error ->", err)
	}
	if user.Email == email {
		err = fmt.Errorf("User Exist")
	}

	return err
}
