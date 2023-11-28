package main

import (
	"fmt"

	"github.com/brandon-a-pinto/nebula/logger-service/configs"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/main/grpc"
	"github.com/brandon-a-pinto/nebula/logger-service/pkg/infra/database"
)

func main() {
	config := configs.LoadConfig()

	// Database
	database.Start(config.DBHost, config.DBName)

	// gRPC Server
	grpc := grpc.NewGRPCServer(config.GRPCServerPort)
	fmt.Println("Starting grpc server on port", config.GRPCServerPort)
	grpc.Start()
}
