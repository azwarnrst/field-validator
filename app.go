package main

import (
	"github.com/azwarnrst/field-validator/internal/router"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	xRouter := router.XRouter{
		FormValidator: router.FormValidator{},
	}
	r := mux.NewRouter()
	r.HandleFunc("/", xRouter.UserHandler2).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}