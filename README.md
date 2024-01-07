# Cocktail Recipe
Welcome to the Cocktail Recipe API project. 
This is a Go project designed for storing and managing cocktail recipes. (I HOPE!)

## Installation

```bash
make build
```

## Run

```bash
make run
```

## JWT Token

To generate new token run `generateToken.go` file manually. Result should be something like this

```
Your Token:
eyJhbGciOiJIUzIJNiIsInR5cCI6IkpXVCJ9.eyJhdRRob3JpemVkIjp0cnVlLCJleHAiOjE3MDUyNTE5ODgsInVzZXJfaWQiOtF3MDQ2NDY5ODg5OTQqDDk1NTl9.Jlhj2LvLC064ZHZbiVCk5b56oRiuaXTuEO2E8i0NkdI
```

Put `Authorization` as a key to the Headers and paste the token with `Bearer ` prefix to access api.