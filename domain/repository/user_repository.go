package repository

import (
	"context"
	"mt-api-go/domain/model"
)

type IUserRepository interface {
	selectUserById(ctx context.Context) (model.MUser, error)
	insertUser(ctx context.Context)
	updateUser(ctx context.Context)
}
