package presenters

import (
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
	"net/http"
)

type MusclePartPresenter struct {
	ctx *gin.Context
}

func NewMusclePartOutputPort(context *gin.Context) ports.MusclePartOutputPort {
	return &MusclePartPresenter{
		ctx: context,
	}
}

func (m *MusclePartPresenter) OutputMuscleParts(muscleParts []*model.MusclePart) error {
	m.ctx.JSON(http.StatusOK, muscleParts)
	return nil
}

func (m *MusclePartPresenter) OutputError(err error) error {
	m.ctx.JSON(http.StatusInternalServerError, err)
	return err
}
