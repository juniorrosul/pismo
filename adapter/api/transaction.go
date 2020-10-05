package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"madsonjr.com/pismo/adapter/serializer"
	"madsonjr.com/pismo/transaction"
)

type transactionHandler struct {
	repository transaction.Repository
	serializer serializer.Transaction
}

// TransactionHandler initializer
func TransactionHandler() *transactionHandler {
	return &transactionHandler{}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	return
}
