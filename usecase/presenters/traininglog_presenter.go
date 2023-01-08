package presenters

import (
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
	"net/http"
)

type TrainingLogPresenter struct {
	ctx *gin.Context
}

func NewTrainingLogOutputPort(context *gin.Context) ports.TrainingLogOutputPort {
	return &TrainingLogPresenter{
		ctx: context,
	}
}

func (l *TrainingLogPresenter) OutputMuscleParts(muscleParts []*model.MusclePart) error {
	l.ctx.JSON(http.StatusOK, muscleParts)
	return nil
}

func (l *TrainingLogPresenter) OutputError(err error) error {
	l.ctx.JSON(http.StatusInternalServerError, err)
	return err
}
