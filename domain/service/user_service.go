package service

import (
	"context"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
)

type IUserService interface {
	FindUserById(ctx context.Context, id string) (*model.MUser, error)
	InsertNewUser(ctx context.Context, user *model.MUser) error
}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUserService {
	return &userService{
		repo: ur,
	}
}

func (us *userService) FindUserById(ctx context.Context, id string) (*model.MUser, error) {
	return us.repo.SelectUserById(ctx, id)
}
func (us *userService) InsertNewUser(ctx context.Context, user *model.MUser) error {
	return us.repo.InsertNewUser(ctx, user)
}
