package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	uniqueId := generateUniqueID()
	token, _ := createToken(uniqueId)

	fmt.Println("Your Token:")
	fmt.Println(token)
}

func generateUniqueID() uint64 {
	timestamp := uint64(time.Now().UnixNano())
	randomNumber := uint64(rand.Int63n(1000000))
	return timestamp + randomNumber
}

func createToken(userID uint64) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")

	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Minute * 60 * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
