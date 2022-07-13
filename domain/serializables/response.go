package serializables

import (
	"fmt"
	"time"

	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/utils"
)

type UserResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt, omitempty"`
}

type AuthUserLoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type ProductResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description, omitempty"`
	Rating      int8   `json:"rating"`
}

type PayloadJWT struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (r *AuthUserLoginResponse) GenerateTokenPayload(mu *models.User) {
	payload := &PayloadJWT{
		Id:        mu.Id,
		Email:     mu.Email,
		Name:      mu.Name,
		CreatedAt: mu.CreateAt,
	}
	token, err := utils.CreateTokenJWT(payload)
	if err != nil {
		fmt.Errorf("Error ->", err)
	}
	r.Token = token
}
