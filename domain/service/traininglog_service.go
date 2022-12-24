package service

import (
	"context"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
)

type ITrainingLogService interface {
	SelectTrainingLogById(ctx context.Context, userId string) ([]*model.TTraininglog, error)
	InsertTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error)
	UpdateTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error)
	DeleteTrainingLog(ctx context.Context, logId int) (int, error)
	DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) (int, string, error)
}

type TrainingLogService struct {
	repo repository.ITrainingLogRepository
}

func NewTrainingLogService(tr repository.ITrainingLogRepository) ITrainingLogService {
	return &TrainingLogService{
		repo: tr,
	}
}

func (ts *TrainingLogService) SelectTrainingLogById(ctx context.Context, userId string) ([]*model.TTraininglog, error) {
	return ts.repo.SelectTrainingLogById(ctx, userId)
}

func (ts *TrainingLogService) InsertTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error) {
	return ts.repo.InsertTrainingLog(ctx, target)
}

func (ts *TrainingLogService) UpdateTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error) {
	return ts.repo.UpdateTrainingLog(ctx, target)
}

func (ts *TrainingLogService) DeleteTrainingLog(ctx context.Context, logId int) (int, error) {
	return ts.repo.DeleteTrainingLog(ctx, logId)
}

func (ts *TrainingLogService) DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) (int, string, error) {
	return ts.repo.DeleteTrainingLogByMenuIdAndDate(ctx, menuId, trainingDate)
}
