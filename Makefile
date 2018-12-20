APP_NAME=golang-challenge
PORT=8080

# Docker Tasks

# Build the container
build: ## Build the container
	docker build . --tag $(APP_NAME)

run: ## run container with desired port to expose
	docker run -i -t --rm -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)

stop: ## stop and remove a running container
	docker stop $(APP_NAME); docker rm $(APP_NAME)