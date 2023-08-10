BROKER_BINARY=broker
LOGGER_BINARY=logger
LISTENER_BINARY=listener
USER_BINARY=user
POST_BINARY=post

prepare:
	@echo "Creating .env file..."
	cp .env.example .env
	@echo "Done!"

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Done!"

up_build: build_broker build_logger build_listener build_user build_post
	@echo "Stopping docker-compose..."
	docker-compose down
	@echo "Building and starting Docker images..."
	docker-compose up --build -d
	@echo "Done!"

down:
	@echo "Stopping docker-compose..."
	docker-compose down
	@echo "Done!"

build_broker:
	@echo "Building broker binary..."
	@cd broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_logger:
	@echo "Building logger binary..."
	@cd logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

build_listener:
	@echo "Building listener binary..."
	@cd listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} ./cmd/listener
	@echo "Done!"

build_user:
	@echo "Building user binary..."
	@cd user-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd/api
	@echo "Done!"

build_post:
	@echo "Building post binary..."
	@cd post-service && env GOOS=linux CGO_ENABLED=0 go build -o ${POST_BINARY} ./cmd/api
	@echo "Done!"
