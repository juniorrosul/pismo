package api

import (
	"github.com/juniorrosul/pismo/adapter/repository/sqlite"
)

type operationTypeHandler struct {
	repository sqlite.OperationType
}

// OperationTypeHandler initializer
func OperationTypeHandler() *operationTypeHandler {
	return &operationTypeHandler{}
}

func (h *operationTypeHandler) Initialize() {
	h.repository.Initialize()
}
