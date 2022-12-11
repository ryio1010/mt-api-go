package service

import (
	"context"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
)

type IUserService interface {
	FindUserById(ctx context.Context, id string) (*model.MUser, error)
	InsertNewUser(ctx context.Context, user *model.MUser) error
	UpdateUser(ctx context.Context, user *model.MUser) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUserService {
	return &UserService{
		repo: ur,
	}
}

func (us *UserService) FindUserById(ctx context.Context, id string) (*model.MUser, error) {
	return us.repo.SelectUserById(ctx, id)
}

func (us *UserService) InsertNewUser(ctx context.Context, user *model.MUser) error {
	return us.repo.InsertNewUser(ctx, user)
}

func (us *UserService) UpdateUser(ctx context.Context, user *model.MUser) error {
	return us.repo.UpdateUser(ctx, user)
}
