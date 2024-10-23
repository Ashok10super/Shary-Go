package main

import (
	"fmt"
	"log"
	"net/http"
)
func Startserver() {

	fmt.Println("Server is starting on port 8080")

	err:= http.ListenAndServe(":8080",nil)
	if err!=nil {
		log.Fatal(err)
	}
}

