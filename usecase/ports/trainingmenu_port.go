package ports

import (
	"context"
	"mt-api-go/usecase/model"
)

type TrainingMenuInputPort interface {
	GetMenuByUserId(ctx context.Context, userId string) error
	InsertMenu(ctx context.Context, menuEntity *model.TrainingMenuRequest) error
	DeleteMenu(ctx context.Context, id int) error
}

type TrainingMenuOutputPort interface {
	OutputTrainingMenu(trainingMenu *model.TrainingMenu) error
	OutputTrainingMenus(trainingMenus []*model.TrainingMenu) error
	OutputTrainingMenuId(id int) error
	OutputError(err error) error
}
