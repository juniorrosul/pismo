package main

import (
	"fmt"
	"net/http"

	"madsonjr.com/pismo/adapter/api"
)

func main() {
	fmt.Println("PISMO")
	accountsHandler := api.NewAccountsHandler()

	http.HandleFunc("/accounts", accountsHandler.Accounts)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
