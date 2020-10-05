package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"madsonjr.com/pismo/adapter/api"
)

func main() {
	fmt.Println("PISMO")
	accountsHandler := api.NewAccountsHandler()

	router := mux.NewRouter()

	router.HandleFunc("/accounts", accountsHandler.Post).Methods("POST")
	router.HandleFunc("/accounts/{id:[0-9]+}", accountsHandler.Get).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
