package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocontactmanager/app"
	"github.com/gocontactmanager/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication) // Attach JWT auth middleware

	port := os.Getenv("PORT") //Get port from .env file
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
