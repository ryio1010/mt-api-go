package interactors

import (
	"context"
	"mt-api-go/domain/repository"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
)

type TrainingLogUseCase struct {
	op   ports.TrainingLogOutputPort
	repo repository.ITrainingLogRepository
}

func NewTrainingLogUseCase(lop ports.TrainingLogOutputPort, lr repository.ITrainingLogRepository) ports.TrainingLogInputPort {
	return &TrainingLogUseCase{
		op:   lop,
		repo: lr,
	}
}

func (l *TrainingLogUseCase) SelectTrainingLogById(ctx context.Context, userId string) error {
	// ログインユーザーの取得
	trainingLogList, err := l.repo.SelectTrainingLogById(ctx, userId)

	if err != nil {
		return err
	}

	response := make([]*model.TrainingLog, 0)
	for _, v := range trainingLogList {
		entity := model.TrainingLogFromDomainModel(v)
		response = append(response, entity)
	}

	return l.op.OutputTrainingLogs(response)
}

func (l *TrainingLogUseCase) InsertTrainingLog(ctx context.Context, target *model.TrainingLog) error {
	return nil
}

func (l *TrainingLogUseCase) UpdateTrainingLog(ctx context.Context, target *model.TrainingLog) error {
	return nil
}

func (l *TrainingLogUseCase) DeleteTrainingLog(ctx context.Context, logId int) error {
	return nil
}

func (l *TrainingLogUseCase) DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) error {
	return nil
}
