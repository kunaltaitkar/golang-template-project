package main

import (
	"net/http"

	"github.com/kunaltaitkar/golang-template-project/config"
	"github.com/kunaltaitkar/golang-template-project/route"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {

	config, err := config.Load(".")
	if err != nil {
		log.Fatalf("failed to load the configuration - %v", err)
		return
	}

	log.SetFormatter(&log.JSONFormatter{})

	router := mux.NewRouter()
	route.Init(router)

	log.Info("server started on ", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
