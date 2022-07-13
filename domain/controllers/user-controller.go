package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/92Sam/ms-users/domain/serializables"
	"github.com/92Sam/ms-users/domain/services"
	"github.com/92Sam/ms-users/utils"
	"github.com/gorilla/mux"
)

type UserController struct {
	*services.Services
}

func NewUserController(svc *services.Services) *UserController {
	return &UserController{svc}
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

func (a *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	resp, err := a.Services.UserService.GetUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	dtoResponseList := make([]*serializables.UserResponse, len(resp))
	for k, v := range resp {
		dtoResponseList[k] = &serializables.UserResponse{
			Id:        v.Id,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreateAt,
		}
	}

	utils.RespondWithJSON(w, http.StatusOK, dtoResponseList)
	return
}

func (a *UserController) GetUsersById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	resp, err := a.Services.UserService.GetUsersById(userId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	if resp == nil {
		utils.RespondWithError(w, http.StatusNoContent, nil)
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

func (a *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	resp, err := a.Services.UserService.DeleteUserById(userId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, resp)
	return
}

func (a *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	p := &serializables.UserRequest{}
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := a.Services.UserService.UpdateUserById(userId, p)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
	return
}
