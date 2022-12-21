package repository

import (
	"context"
	"mt-api-go/domain/model"
)

type ITrainingMenuRepository interface {
	SelectMenuByUserId(ctx context.Context, userId string) ([]*model.MMenu, error)
	InsertMenu(ctx context.Context, menuEntity *model.MMenu) (*model.MMenu, error)
	DeleteMenu(ctx context.Context, id int) (int, error)
}
