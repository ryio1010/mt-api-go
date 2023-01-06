package interactors

import (
	"context"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
)

type ITrainingLogUseCase interface {
	SelectTrainingLogById(ctx context.Context, userId string) ([]*model.TrainingLog, error)
	InsertTrainingLog(ctx context.Context, target *model.TrainingLog) (*model.TrainingLog, error)
	UpdateTrainingLog(ctx context.Context, target *model.TrainingLog) (*model.TrainingLog, error)
	DeleteTrainingLog(ctx context.Context, logId int) (int, error)
	DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) (int, string, error)
}

type TrainingLogUseCase struct {
	svc service.ITrainingLogService
}

func NewTrainingLogUseCase(ls service.ITrainingLogService) ITrainingLogUseCase {
	return &TrainingLogUseCase{
		svc: ls,
	}
}

func (lu *TrainingLogUseCase) SelectTrainingLogById(ctx context.Context, userId string) ([]*model.TrainingLog, error) {
	// ログインユーザーの取得
	trainingLogList, err := lu.svc.SelectTrainingLogById(ctx, userId)

	if err != nil {
		return nil, err
	}

	response := make([]*model.TrainingLog, 0)
	for _, v := range trainingLogList {
		entity := model.TrainingLogFromDomainModel(v)
		response = append(response, entity)
	}

	return response, nil
}

func (lu *TrainingLogUseCase) InsertTrainingLog(ctx context.Context, target *model.TrainingLog) (*model.TrainingLog, error) {
	return nil, nil
}

func (lu *TrainingLogUseCase) UpdateTrainingLog(ctx context.Context, target *model.TrainingLog) (*model.TrainingLog, error) {
	return nil, nil
}

func (lu *TrainingLogUseCase) DeleteTrainingLog(ctx context.Context, logId int) (int, error) {
	return 0, nil
}

func (lu *TrainingLogUseCase) DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) (int, string, error) {
	return 0, "", nil
}
