package repository

import (
	"context"
	"mt-api-go/domain/model"
)

type IUserRepository interface {
	SelectUserById(ctx context.Context, id string) (*model.MUser, error)
	InsertNewUser(ctx context.Context, target *model.MUser) error
}
