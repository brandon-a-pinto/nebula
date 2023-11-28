BROKER_BINARY=broker
LOGGER_BINARY=logger
LISTENER_BINARY=listener
USER_BINARY=user
POST_BINARY=post

prepare:
	@echo "Generating .env file..."
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
	@cd broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/bin/${BROKER_BINARY} ./cmd/broker
	@echo "Done!"

build_logger:
	@echo "Building logger binary..."
	@cd logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/bin/${LOGGER_BINARY} ./cmd/logger
	@echo "Done!"

build_listener:
	@echo "Building listener binary..."
	@cd listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/bin/${LISTENER_BINARY} ./cmd/listener
	@echo "Done!"

build_user:
	@echo "Building user binary..."
	@cd user-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/bin/${USER_BINARY} ./cmd/user
	@echo "Done!"

build_post:
	@echo "Building post binary..."
	@cd post-service && env GOOS=linux CGO_ENABLED=0 go build -o ./build/bin/${POST_BINARY} ./cmd/post
	@echo "Done!"

grpc:
	@echo "Generating gRPC files..."
	@cd logger-service && protoc --go_out=. --go-grpc_out=. ./internal/main/grpc/protofile/*.proto
	@cd user-service && protoc --go_out=. --go-grpc_out=. ./internal/main/grpc/protofile/*.proto
	@cd broker-service && protoc --go_out=. --go-grpc_out=. ./internal/main/grpc/protofile/*.proto
	@echo "Done!"
