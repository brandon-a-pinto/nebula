package main

import (
	"fmt"

	"github.com/brandon-a-pinto/nebula/user-service/configs"
	"github.com/brandon-a-pinto/nebula/user-service/internal/main/grpc"
)

func main() {
	config := configs.LoadConfig()

	// gRPC Server
	grpc := grpc.NewGRPCServer(config.GRPCServerPort)
	fmt.Println("Starting grpc server on port", config.GRPCServerPort)
	grpc.Start()
}
