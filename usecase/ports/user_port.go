package ports

import (
	"context"
	"mt-api-go/usecase/model"
)

type UserInputPort interface {
	LoginUser(ctx context.Context, user *model.User) error
	InsertNewUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
}

type UserOutputPort interface {
	OutputUser(user *model.User) error
	OutputError(err error) error
}
