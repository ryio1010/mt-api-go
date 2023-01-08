package presenters

import (
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
	"net/http"
)

type TrainingMenuPresenter struct {
	ctx *gin.Context
}

func NewTrainingMenuOutputPort(context *gin.Context) ports.TrainingMenuOutputPort {
	return &TrainingMenuPresenter{
		ctx: context,
	}
}

func (m *TrainingMenuPresenter) OutputTrainingMenu(trainingMenu *model.TrainingMenu) error {
	m.ctx.JSON(http.StatusOK, trainingMenu)
	return nil
}

func (m *TrainingMenuPresenter) OutputTrainingMenus(trainingMenus []*model.TrainingMenu) error {
	m.ctx.JSON(http.StatusOK, trainingMenus)
	return nil
}

func (m *TrainingMenuPresenter) OutputError(err error) error {
	m.ctx.JSON(http.StatusInternalServerError, err)
	return err
}
