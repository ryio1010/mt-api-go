package repository

import (
	"context"
	"mt-api-go/domain/model"
)

type ITrainingLogRepository interface {
	SelectTrainingLogById(ctx context.Context, userId string) ([]*model.TTraininglog, error)
	InsertTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error)
	UpdateTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error)
	DeleteTrainingLog(ctx context.Context, logId int) (int, error)
	DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) (int, string, error)
}
