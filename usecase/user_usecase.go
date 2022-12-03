package usecase

import (
	"context"
	models "mt-api-go/domain/model"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
)

type IUserUseCase interface {
	FindUserById(ctx context.Context, id string) (*model.User, error)
	InsertNewUser(ctx context.Context, user *model.User) error
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

func (uu *userUseCase) InsertNewUser(ctx context.Context, user *model.User) error {
	insertionTarget := models.MUser{
		Userid:   string(user.ID),
		Username: user.Name,
		Password: user.Password,
	}
	err := uu.svc.InsertNewUser(ctx, &insertionTarget)

	if err != nil {
		return err
	}

	return nil
}
