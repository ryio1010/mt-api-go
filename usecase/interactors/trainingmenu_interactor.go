package interactors

import (
	"context"
	"github.com/volatiletech/null/v8"
	models "mt-api-go/domain/model"
	"mt-api-go/domain/repository"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
)

type TrainingMenuUseCase struct {
	op   ports.TrainingMenuOutputPort
	repo repository.ITrainingMenuRepository
}

func NewTrainingMenuUseCase(mop ports.TrainingMenuOutputPort, mr repository.ITrainingMenuRepository) ports.TrainingMenuInputPort {
	return &TrainingMenuUseCase{
		op:   mop,
		repo: mr,
	}
}

func (m *TrainingMenuUseCase) GetMenuByUserId(ctx context.Context, userId string) error {
	// ログインユーザーの取得
	trainingMenus, err := m.repo.SelectMenuByUserId(ctx, userId)

	if err != nil {
		return err
	}

	response := make([]*model.TrainingMenu, 0)
	for _, v := range trainingMenus {
		entity := model.TrainingMenuFromDomainModel(v)
		response = append(response, entity)
	}

	return m.op.OutputTrainingMenus(response)
}

func (m *TrainingMenuUseCase) InsertMenu(ctx context.Context, menuEntity *model.TrainingMenuRequest) error {

	insertionTarget := models.MMenu{
		Menuname:     menuEntity.MenuName,
		Musclepartid: menuEntity.MusclePartId,
		Userid:       null.StringFrom(menuEntity.UserId),
		Status:       "1",
	}
	result, err := m.repo.InsertMenu(ctx, &insertionTarget)

	if err != nil {
		return err
	}

	return m.op.OutputTrainingMenu(model.TrainingMenuFromDomainModel(result))
}

func (m *TrainingMenuUseCase) DeleteMenu(ctx context.Context, id int) error {
	deletedId, err := m.repo.DeleteMenu(ctx, id)

	if err != nil {
		return err
	}

	return m.op.OutputTrainingMenuId(deletedId)
}
