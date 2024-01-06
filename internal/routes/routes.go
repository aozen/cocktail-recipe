package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aozen/cocktail-recipe/internal/models"
)

var collection *mongo.Collection

func SetupRoutes(r *mux.Router, db *mongo.Database) {
	collection = db.Collection("cocktails")

	r.HandleFunc("/api/cocktails", getCocktails).Methods("GET")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getCocktails(w http.ResponseWriter, _ *http.Request) {
	var cocktails []models.Cocktail
	var cocktail models.Cocktail

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Error fetching cocktails: %v", err)
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error fetching cocktails"})
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		if err := cur.Decode(&cocktail); err != nil {
			log.Printf("Error decoding cocktail: %v", err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error decoding cocktails"})
			return
		}
		cocktails = append(cocktails, cocktail)
	}

	respondWithJSON(w, http.StatusOK, cocktails)
}
