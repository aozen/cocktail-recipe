package main

import (
	"fmt"
	"github.com/aozen/cocktail-recipe/internal/database"
	"github.com/aozen/cocktail-recipe/internal/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	mongoClient, err := database.InitMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(nil) //In localhost, I am closing with Ctrl + C so looks like it's not useful right now.

	router := mux.NewRouter()
	routes.SetupRoutes(router, mongoClient.Database("cocktail"))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server stopped")
}
