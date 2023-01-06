package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase/interactors"
	"net/http"
)

type MusclePartHandler struct {
	usecase interactors.IMusclePartUseCase
}

func NewMusclePartHandler(mu interactors.IMusclePartUseCase) *MusclePartHandler {
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
