package grpc

import (
	"fmt"
	"net"

	"github.com/brandon-a-pinto/nebula/logger-service/internal/main/factory"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/main/grpc/pb"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/presentation/service"
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

func (s *GRPCServer) services(server grpc.ServiceRegistrar) {
	logService := service.NewLogService(
		*factory.CreateLogFactory(),
	)

	pb.RegisterLogServiceServer(server, logService)
}

func (s *GRPCServer) Start() {
	server := grpc.NewServer()

	s.services(server)
	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
