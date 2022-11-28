package usecase

import (
	"context"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
)

type IUserUseCase interface {
	FindUserById(ctx context.Context, id string) (*model.User, error)
}

type userUseCase struct {
	svc service.IUserService
}

func NewUserUseCase(us service.IUserService) IUserUseCase {
	return &userUseCase{
		svc: us,
	}
}

func (uu *userUseCase) FindUserById(ctx context.Context, id string) (*model.User, error) {
	ms, err := uu.svc.FindUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return model.UserFromDomainModel(ms), nil
}
