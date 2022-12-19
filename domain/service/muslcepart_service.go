package service

import (
	"context"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
)

type IMusclePartService interface {
	SelectAllMusclePart(ctx context.Context) ([]*model.MMusclepart, error)
}

type MusclePartService struct {
	repo repository.IMusclePartRepository
}

func NewMusclePartService(mr repository.IMusclePartRepository) IMusclePartService {
	return &MusclePartService{
		repo: mr,
	}
}

func (ms *MusclePartService) SelectAllMusclePart(ctx context.Context) ([]*model.MMusclepart, error) {
	return ms.repo.SelectAllMusclePart(ctx)
}
