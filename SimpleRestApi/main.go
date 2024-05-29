package main

import (
	"SimpleRestApi/handler"
	"SimpleRestApi/model"
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Staring Service ")

	model.UserData = []model.User{}

	restHandler := handler.RestServiceHandler()
	http.HandleFunc("/api/add", restHandler.AddUser)
	http.HandleFunc("/api/get", restHandler.GetUser)

	fmt.Println("Running server at :80 ")
	http.ListenAndServe(":80", nil)
}
