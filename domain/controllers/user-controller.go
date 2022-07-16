package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/92Sam/ms-users/domain/serializables"
	"github.com/92Sam/ms-users/domain/services"
	"github.com/92Sam/ms-users/utils"
	"github.com/gorilla/mux"
)

type IUserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUsersById(w http.ResponseWriter, r *http.Request)
	DeleteUserById(w http.ResponseWriter, r *http.Request)
	UpdateUserById(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	*services.Services
}

func NewUserController(svc *services.Services) IAuthUserController {
	return &UserController{svc}
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
