package main

import (
	"fmt"
	"log"
	"net/http"
)
func Startserver(mux *http.ServeMux) {

	fmt.Println("Server is starting on port 8080")

	err:= http.ListenAndServe(":8080",mux)
	if err!=nil {
		log.Fatal(err)
	}
}

