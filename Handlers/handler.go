package handlers

import (
	"net/http"

)

func Homepagehandler(w http.ResponseWriter,r* http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type","Text/plain")
    w.Write([]byte("Hi there Welcome to the Sharpy site"))
}