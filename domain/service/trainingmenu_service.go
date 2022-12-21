package service

import (
	"context"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
)

type ITrainingMenuService interface {
	SelectMenuByUserId(ctx context.Context, userId string) ([]*model.MMenu, error)
	InsertMenu(ctx context.Context, menuEntity *model.MMenu) (*model.MMenu, error)
	DeleteMenu(ctx context.Context, id int) (int, error)
}

type TrainingMenuService struct {
	repo repository.ITrainingMenuRepository
}

func NewTrainingMenuService(tr repository.ITrainingMenuRepository) ITrainingMenuService {
	return &TrainingMenuService{
		repo: tr,
	}
}

func (ts *TrainingMenuService) SelectMenuByUserId(ctx context.Context, userId string) ([]*model.MMenu, error) {
	return ts.repo.SelectMenuByUserId(ctx, userId)
}

func (ts *TrainingMenuService) InsertMenu(ctx context.Context, menuEntity *model.MMenu) (*model.MMenu, error) {
	return ts.repo.InsertMenu(ctx, menuEntity)
}

func (ts *TrainingMenuService) DeleteMenu(ctx context.Context, id int) (int, error) {
	return ts.repo.DeleteMenu(ctx, id)
}
