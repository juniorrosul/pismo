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
	fmt.Println("PISMO - API")

	router := mux.NewRouter()

	// Operation Type
	api.OperationTypeHandler().Initialize()

	// Accounts
	accountsHandler := api.AccountsHandler()
	accountsHandler.Initialize()
	router.HandleFunc("/accounts", accountsHandler.AccountPost).Methods("POST")
	router.HandleFunc("/accounts/{id:[0-9]+}", accountsHandler.AccountGet).Methods("GET")

	// Transactions
	transactionHandler := api.TransactionHandler()
	transactionHandler.Initialize()
	router.HandleFunc("/transactions", transactionHandler.TransactionPost).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
