package api

import (
	"github.com/juniorrosul/pismo/adapter/repository/mysqlconnection"
)

type operationTypeHandler struct {
	repository mysqlconnection.OperationType
}

// OperationTypeHandler initializer
func OperationTypeHandler() *operationTypeHandler {
	return &operationTypeHandler{}
}

func (h *operationTypeHandler) Initialize() {
	h.repository.Initialize()
}
