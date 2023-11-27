package main

import (
	"fmt"

	"github.com/brandon-a-pinto/nebula/post-service/configs"
	"github.com/brandon-a-pinto/nebula/post-service/internal/main/grpc"
)

func main() {
	config := configs.LoadConfig()

	// gRPC Server
	grpc := grpc.NewGRPCServer(config.GRPCServerPort)
	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	grpc.Start()
}
