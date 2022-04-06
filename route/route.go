package route

import (
	"net/http"

	"github.com/kunaltaitkar/golang-template-project/handler"
	"github.com/kunaltaitkar/golang-template-project/model"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {

	var factory model.Factory = FactoryImpl{}

	router.HandleFunc("/user/{userId}", handler.User(factory).Get).Methods(http.MethodGet)
	router.HandleFunc("/user", handler.User(factory).Post).Methods(http.MethodPost)

}
