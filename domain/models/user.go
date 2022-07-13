package models

import (
	"fmt"
	"time"

	"github.com/92Sam/ms-users/utils"
)

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreateAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt, omitempty"`
}

func (u *User) CompareHashAndPassword(passwordUncrypted string) error {
	if !utils.CheckPasswordHash(passwordUncrypted, u.Password) {
		return fmt.Errorf("password or email invalid")
	}
	return nil
}
