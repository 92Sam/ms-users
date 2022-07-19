package serializables

import (
	"fmt"

	"github.com/92Sam/ms-users/utils"
)

type UserRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description", omitempty`
	Rating      int8    `json:"rating", omitempty`
}

type AuthUserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserSignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ur *UserRequest) PasswordHashed() string {
	passwordHashed, err := utils.HashPassword(*ur.Password)
	if err != nil {
		fmt.Errorf("Error ->", err)
	}
	return passwordHashed
}

func (ur *AuthUserSignupRequest) PasswordHashed() string {
	passwordHashed, err := utils.HashPassword(ur.Password)
	if err != nil {
		fmt.Errorf("Error ->", err)
	}
	return passwordHashed
}
