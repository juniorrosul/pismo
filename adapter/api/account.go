package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/juniorrosul/pismo/adapter/repository/sqlite"
	"github.com/juniorrosul/pismo/adapter/serializer"
)

type accountHandler struct {
	serializer serializer.Account
	repository sqlite.Account
}

// AccountsHandler handler initializer
func AccountsHandler() *accountHandler {
	return &accountHandler{}
}

func (h *accountHandler) Initialize() {
	h.repository.Initialize()
}

func (h *accountHandler) AccountGet(w http.ResponseWriter, r *http.Request) {
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

func (h *accountHandler) AccountPost(w http.ResponseWriter, r *http.Request) {
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
