package interactors

import (
	"context"
	"mt-api-go/domain/repository"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
)

type IMusclePartUseCase interface {
	GetAllMusclePart(ctx context.Context) ([]*model.MusclePart, error)
}

type MusclePartUseCase struct {
	op   ports.MusclePartOutputPort
	repo repository.IMusclePartRepository
}

func NewMusclePartUseCase(mop ports.MusclePartOutputPort, mr repository.IMusclePartRepository) ports.MusclePartInputPort {
	return &MusclePartUseCase{
		op:   mop,
		repo: mr,
	}
}

func (mu *MusclePartUseCase) GetAllMusclePart(ctx context.Context) error {
	// ログインユーザーの取得
	muscleParts, err := mu.repo.SelectAllMusclePart(ctx)

	if err != nil {
		return err
	}

	response := make([]*model.MusclePart, 0)
	for _, v := range muscleParts {
		entity := model.MusclePartFromDomainModel(v)
		response = append(response, entity)
	}

	return mu.op.OutputMuscleParts(response)
}
