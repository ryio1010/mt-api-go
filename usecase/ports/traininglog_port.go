package ports

import (
	"context"
	"mt-api-go/usecase/model"
)

type TrainingLogInputPort interface {
	SelectTrainingLogById(ctx context.Context, userId string) error
	InsertTrainingLog(ctx context.Context, target *model.TrainingLog) error
	UpdateTrainingLog(ctx context.Context, target *model.TrainingLog) error
	DeleteTrainingLog(ctx context.Context, logId int) error
	DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) error
}

type TrainingLogOutputPort interface {
	OutputTrainingLog(trainingLog *model.TrainingLog) error
	OutputTrainingLogs(trainingLogs []*model.TrainingLog) error
	OutputLogId(logId int) error
	OutputMenuIdAndTrainingDate(menuId int, trainingDate string) error
	OutputError(err error) error
}
