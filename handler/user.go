package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kunaltaitkar/golang-template-project/model"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type userHandler struct {
	factory model.UserServiceFactory
}

func User(factory model.UserServiceFactory) *userHandler {
	return &userHandler{factory: factory}
}

func (u userHandler) Post(response http.ResponseWriter, request *http.Request) {

	var userDetails model.User

	if err := json.NewDecoder(request.Body).Decode(&userDetails); err != nil {
		log.Error("invalid request body: %v", err)
		http.Error(response, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := u.factory.GetUserService().CreateUser(userDetails)
	if err != nil {
		log.Error("failed to create user: %v", err)
		http.Error(response, "failed to create user", http.StatusInternalServerError)
		return
	}

	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(user)
}

func (u userHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	user, err := u.factory.GetUserService().GetUser(userId)
	if err != nil {
		log.Error("failed to get user: %v", err)
		http.Error(response, "failed to get user", http.StatusInternalServerError)
		return
	}
	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(user)
}
