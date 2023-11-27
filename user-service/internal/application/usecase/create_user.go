package usecase

import (
	"context"
	"time"

	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/contract"
	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/dto"
	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/entity"
)

type CreateUserUsecase struct {
	UserRepository contract.IUserRepository
	BcryptAdapter  contract.IBcryptAdapter
}

func NewCreateUserUsecase(userRepository contract.IUserRepository, bcryptAdapter contract.IBcryptAdapter) *CreateUserUsecase {
	return &CreateUserUsecase{
		UserRepository: userRepository,
		BcryptAdapter:  bcryptAdapter,
	}
}

func (u *CreateUserUsecase) Exec(c context.Context, input dto.CreateUserInput) (*dto.CreateUserOutput, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*3)
	defer cancel()

	user, err := entity.NewUser(input)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := u.BcryptAdapter.Hash(user.Password, 12)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = u.UserRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	output := &dto.CreateUserOutput{
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
	}

	return output, nil
}
