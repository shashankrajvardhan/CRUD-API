package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func router(router *mux.Router) {
	//	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/getusers/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/createusers", createUser(db)).Methods("POST")
	router.HandleFunc("/updateusers/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/deleteusers/{id}", deleteUser(db)).Methods("DELETE")

	router.Use(jsonContentTypeMiddleware)
	//start server
	fmt.Println("Running in Port:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
	//	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}
