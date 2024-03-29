package routes

import (
	"context"
	"encoding/json"
	"github.com/aozen/cocktail-recipe/internal/middleware"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aozen/cocktail-recipe/internal/models"
)

var collection *mongo.Collection

func SecureAPI(handler http.HandlerFunc) http.HandlerFunc {
	return middleware.TokenVerifyMiddleware(handler)
}

func SetupRoutes(r *mux.Router, db *mongo.Database) {
	collection = db.Collection("cocktails")

	r.HandleFunc("/api/cocktails", SecureAPI(getCocktails)).Methods("GET")
	r.HandleFunc("/api/cocktails/{id}", SecureAPI(getCocktail)).Methods("GET")
	r.HandleFunc("/api/cocktails", SecureAPI(createCocktail)).Methods("POST")
	r.HandleFunc("/api/cocktails/{id}", SecureAPI(updateCocktail)).Methods("PUT")
	r.HandleFunc("/api/cocktails/{id}", SecureAPI(deleteCocktail)).Methods("DELETE")
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

func getCocktail(w http.ResponseWriter, r *http.Request) {
	var cocktail models.Cocktail

	params := mux.Vars(r)
	cocktailID := params["id"]

	err := collection.FindOne(context.Background(), bson.D{{"_id", cocktailID}}).Decode(&cocktail)

	if err != nil {
		log.Printf("Error fetching cocktail: %v", err)
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error fetching cocktail"})
		return
	}

	respondWithJSON(w, http.StatusOK, cocktail)
}

func createCocktail(w http.ResponseWriter, r *http.Request) {
	var newCocktail models.Cocktail

	_ = json.NewDecoder(r.Body).Decode(&newCocktail)
	id := uuid.New()
	newCocktail.ID = id.String()

	_, err := collection.InsertOne(context.Background(), newCocktail)
	if err != nil {
		log.Printf("Error creating cocktail: %v", err)
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error creating cocktail"})
		return
	}

	respondWithJSON(w, http.StatusCreated, newCocktail)
}

// Instead of update it's a replacement right now. Doesn't sure if its worth for checking every prop and updating matches
// TODO: Check&Rewrite after a couple of months if this is not forgotten
func updateCocktail(w http.ResponseWriter, r *http.Request) {
	var replacementCocktail models.Cocktail

	params := mux.Vars(r)
	cocktailID := params["id"]

	_ = json.NewDecoder(r.Body).Decode(&replacementCocktail)

	_, err := collection.ReplaceOne(context.Background(), bson.D{{"_id", cocktailID}}, replacementCocktail)
	if err != nil {
		log.Printf("Error updating cocktail: %v", err)
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error updating cocktail"})
		return
	}

	respondWithJSON(w, http.StatusOK, replacementCocktail)
}

func deleteCocktail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cocktailID := params["id"]

	_, err := collection.DeleteOne(context.Background(), bson.D{{"_id", cocktailID}})
	if err != nil {
		log.Printf("Error deleting cocktail: %v", err)
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error deleting cocktail"})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Removed cocktail"})
}
