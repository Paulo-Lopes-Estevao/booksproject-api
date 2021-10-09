package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/bookproject/controllers"
	"github.com/gorilla/mux"
)

var Version string = "/v1"

func main() {
	route := mux.NewRouter()

	route.HandleFunc(Version+"/category", controllers.ListCategoryAdapter).Methods("GET")
	route.HandleFunc(Version+"/category", controllers.CreateCategoryAdapter).Methods("POST")

	fmt.Println("Server started at port 2000")
	err := http.ListenAndServe(":2000", route)

	if err != nil {
		log.Println("Not Running Server...", err.Error())
	}
}
