package interactors

import (
	"context"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
)

type IMusclePartUseCase interface {
	GetAllMusclePart(ctx context.Context) ([]*model.MusclePart, error)
}

type MusclePartUseCase struct {
	svc service.IMusclePartService
}

func NewMusclePartUseCase(ms service.IMusclePartService) IMusclePartUseCase {
	return &MusclePartUseCase{
		svc: ms,
	}
}

func (mu *MusclePartUseCase) GetAllMusclePart(ctx context.Context) ([]*model.MusclePart, error) {
	// ログインユーザーの取得
	muscleParts, err := mu.svc.SelectAllMusclePart(ctx)

	if err != nil {
		return nil, err
	}

	response := make([]*model.MusclePart, 0)
	for _, v := range muscleParts {
		entity := model.MusclePartFromDomainModel(v)
		response = append(response, entity)
	}

	return response, nil
}
