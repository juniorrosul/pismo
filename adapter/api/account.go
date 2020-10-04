package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"madsonjr.com/pismo/account"
	"madsonjr.com/pismo/adapter/serializer"
)

type handler struct {
	service    account.Service
	serializer serializer.Account
}

func NewAccountsHandler() *handler {
	return &handler{}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
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

	_, err = h.serializer.Decode(bodyBytes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid content")))
		return
	}

	// err = h.service.Store(account)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(fmt.Sprintf("Invalid content")))
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	return
}

func (h *handler) Accounts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.Post(w, r)
		return
	}
}
