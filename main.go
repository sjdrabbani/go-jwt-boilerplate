package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocontactmanager/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
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
