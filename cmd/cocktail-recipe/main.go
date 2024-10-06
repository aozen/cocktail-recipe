package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aozen/cocktail-recipe/internal/database"
	"github.com/aozen/cocktail-recipe/internal/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	mongoClient, err := database.InitMongoDB("mongodb://mongo:27017")
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
