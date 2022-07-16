package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/92Sam/ms-users/domain/serializables"
	"github.com/92Sam/ms-users/utils"
)

type IAuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
}

func (a *UserController) Login(w http.ResponseWriter, r *http.Request) {
	p := &serializables.AuthUserLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.Services.AuthService.Login(p)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	resp := &serializables.AuthUserLoginResponse{
		Id:    user.Email,
		Email: user.Email,
		Name:  user.Name,
	}
	resp.GenerateTokenPayload(user)

	utils.RespondWithJSON(w, http.StatusOK, resp)
	return
}

func (a *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	p := &serializables.AuthUserSignupRequest{}
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := a.Services.AuthService.Signup(p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	dtoResponse := &serializables.UserResponse{
		Id:        resp.Id,
		Name:      resp.Name,
		Email:     resp.Email,
		CreatedAt: resp.CreateAt,
	}

	utils.RespondWithJSON(w, http.StatusCreated, dtoResponse)
	return
}
