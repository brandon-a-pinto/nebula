package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Port string
}

func NewGRPCServer(port string) *GRPCServer {
	return &GRPCServer{
		Port: port,
	}
}

func (s *GRPCServer) Start() {
	server := grpc.NewServer()

	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
