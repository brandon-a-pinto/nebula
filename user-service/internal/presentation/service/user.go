package service

import (
	"context"

	"github.com/brandon-a-pinto/nebula/user-service/internal/application/usecase"
	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/dto"
	"github.com/brandon-a-pinto/nebula/user-service/internal/main/grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	CreateUserUsecase usecase.CreateUserUsecase
}

func NewUserService(createUserUsecase usecase.CreateUserUsecase) *UserService {
	return &UserService{
		CreateUserUsecase: createUserUsecase,
	}
}

func (s *UserService) CreateUser(ctx context.Context, input *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	dto := dto.CreateUserInput{
		Email:       input.Email,
		Username:    input.Username,
		DisplayName: input.DisplayName,
		Password:    input.Password,
	}

	output, err := s.CreateUserUsecase.Exec(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Email:       output.Email,
		Username:    output.Username,
		DisplayName: output.DisplayName,
	}, nil
}
