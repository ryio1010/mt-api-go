package usecase

import (
	"context"
	"github.com/volatiletech/null/v8"
	models "mt-api-go/domain/model"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
)

type ITrainingMenuUseCase interface {
	GetMenuByUserId(ctx context.Context, userId string) ([]*model.TrainingMenu, error)
	InsertMenu(ctx context.Context, menuEntity *model.TrainingMenuRequest) (*model.TrainingMenu, error)
	DeleteMenu(ctx context.Context, id int) (int, error)
}

type TrainingMenuUseCase struct {
	svc service.ITrainingMenuService
}

func NewTrainingMenuUseCase(ms service.ITrainingMenuService) ITrainingMenuUseCase {
	return &TrainingMenuUseCase{
		svc: ms,
	}
}

func (mu *TrainingMenuUseCase) GetMenuByUserId(ctx context.Context, userId string) ([]*model.TrainingMenu, error) {
	// ログインユーザーの取得
	trainingMenus, err := mu.svc.SelectMenuByUserId(ctx, userId)

	if err != nil {
		return nil, err
	}

	response := make([]*model.TrainingMenu, 0)
	for _, v := range trainingMenus {
		entity := model.TrainingMenuFromDomainModel(v)
		response = append(response, entity)
	}

	return response, nil
}

func (mu *TrainingMenuUseCase) InsertMenu(ctx context.Context, menuEntity *model.TrainingMenuRequest) (*model.TrainingMenu, error) {

	insertionTarget := models.MMenu{
		Menuname:     menuEntity.MenuName,
		Musclepartid: menuEntity.MusclePartId,
		Userid:       null.StringFrom(menuEntity.UserId),
		Status:       "1",
	}
	result, err := mu.svc.InsertMenu(ctx, &insertionTarget)

	if err != nil {
		return nil, err
	}

	return model.TrainingMenuFromDomainModel(result), nil
}

func (mu *TrainingMenuUseCase) DeleteMenu(ctx context.Context, id int) (int, error) {
	deletedId, err := mu.svc.DeleteMenu(ctx, id)

	if err != nil {
		return 0, err
	}

	return deletedId, nil
}
