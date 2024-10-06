IMAGE_NAME := cocktail-recipe-app
CONTAINER_NAME := cocktail-recipe-container
MONGO_NAME := mongo
PORT := 8080
NETWORK := cocktail-network

build:
	docker build -t $(IMAGE_NAME) .

network:
	@docker network inspect $(NETWORK) >/dev/null 2>&1 || docker network create $(NETWORK)

mongo: network
	@docker run -d --name $(MONGO_NAME) --network $(NETWORK) mongo || echo "Mongo container already running."

run: network
	docker run -d --name $(CONTAINER_NAME) --network $(NETWORK) -p $(PORT):$(PORT) $(IMAGE_NAME)

stop:
	@docker stop $(CONTAINER_NAME) || true
	@docker rm $(CONTAINER_NAME) || true
	@docker stop $(MONGO_NAME) || true
	@docker rm $(MONGO_NAME) || true

clean:
	@docker network rm $(NETWORK) || true

all: build mongo run

refresh: stop clean all
