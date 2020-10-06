package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juniorrosul/pismo/adapter/repository/mysqlconnection"
	"github.com/juniorrosul/pismo/adapter/serializer"
)

type transactionHandler struct {
	repository mysqlconnection.Transaction
	serializer serializer.Transaction
}

// TransactionHandler initializer
func TransactionHandler() *transactionHandler {
	return &transactionHandler{}
}

func (h *transactionHandler) Initialize() {
	h.repository.Initialize()
}

func (h *transactionHandler) TransactionPost(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Invalid content type: %s", ct)))
		return
	}

	transaction, err := h.serializer.Decode(bodyBytes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid content")))
		panic(err.Error())
	}
	fmt.Println(transaction)

	err = h.repository.Store(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error")))
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return
}
