FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin cmd/cocktail-recipe/main.go

EXPOSE 8080

CMD ["./bin/main"]