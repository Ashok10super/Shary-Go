package routes

import (
	"net/http"
	handlers "shary/Handlers"
)


func Setroutes()*http.ServeMux  {//This http.serveMux intialies the routes
	
mux:=http.NewServeMux(); //creates the new servemax

//diffining the routes using mux struct and  methods associate inside the mux

mux.HandleFunc("/",handlers.Homepagehandler)

return mux

}