package main

import (
	"fmt"
	routes "shary/Routes"
)
func main()  {
	fmt.Println("System started")

	mux:= routes.Setroutes()

	Startserver(mux)

}