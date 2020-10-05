package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	repository "madsonjr.com/pismo/adapter/repository/sqlite"
	"madsonjr.com/pismo/adapter/serializer"
)

type handler struct {
	serializer serializer.Account
	repository repository.Account
}

// NewAccountsHandler handler initializer
func NewAccountsHandler() *handler {
	return &handler{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	urlParser := strings.Split(r.URL.String(), "/")
	if len(urlParser) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountID, err := strconv.Atoi(urlParser[2])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	account, err := h.repository.Find(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	encoded, err := h.serializer.Encode(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
	return
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

	account, err := h.serializer.Decode(bodyBytes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid content")))
		return
	}

	err = h.repository.Store(account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid content")))
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
