package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase"
	"net/http"
)

type MusclePartHandler struct {
	usecase usecase.IMusclePartUseCase
}

func NewMusclePartHandler(mu usecase.IMusclePartUseCase) *MusclePartHandler {
	return &MusclePartHandler{
		usecase: mu,
	}
}

func (mh *MusclePartHandler) GetAllMusclePart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		muscleParts, err := mh.usecase.GetAllMusclePart(ctx)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, muscleParts)
	}
}
