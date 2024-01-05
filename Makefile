IMAGE_NAME := cocktail-recipe-app
CONTAINER_NAME := cocktail-recipe-container
PORT := 8080

build:
	docker build -t $(IMAGE_NAME) .

# PROD: I may want to add unique name ===> --name $(CONTAINER_NAME)
run:
	docker run -p $(PORT):$(PORT) $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)