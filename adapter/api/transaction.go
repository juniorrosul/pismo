package api

import (
	"net/http"

	repository "madsonjr.com/pismo/adapter/repository/sqlite"
)

type transactionHandler struct {
	repository repository.OperationType
}

// TransactionHandler initializer
func TransactionHandler() *transactionHandler {
	return &transactionHandler{}
}

func (h *transactionHandler) TransactionPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}
