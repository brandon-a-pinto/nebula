package usecase

import (
	"context"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/dto"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/main/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateUserUsecase struct {
	UserGRPCPort string
}

func NewCreateUserUsecase(port string) *CreateUserUsecase {
	return &CreateUserUsecase{
		UserGRPCPort: port,
	}
}

func (u *CreateUserUsecase) Exec(ctx context.Context, input dto.CreateUserInput) (*pb.CreateUserResponse, error) {
	conn, err := grpc.Dial("user-service:"+u.UserGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	output, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Email:       input.Email,
		Username:    input.Username,
		DisplayName: input.DisplayName,
		Password:    input.Password,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}
