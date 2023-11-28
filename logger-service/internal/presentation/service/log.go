package service

import (
	"context"

	"github.com/brandon-a-pinto/nebula/logger-service/internal/application/usecase"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/domain/dto"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/main/grpc/pb"
)

type LogService struct {
	pb.UnimplementedLogServiceServer
	CreateLogUsecase usecase.CreateLogUsecase
}

func NewLogService(createLogUsecase usecase.CreateLogUsecase) *LogService {
	return &LogService{
		CreateLogUsecase: createLogUsecase,
	}
}

func (s *LogService) CreateLog(ctx context.Context, input *pb.CreateLogRequest) (*pb.CreateLogResponse, error) {
	dto := dto.CreateLogInput{
		Msg:  input.Msg,
		Type: input.Type,
	}

	err := s.CreateLogUsecase.Exec(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateLogResponse{}, nil
}
