package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase/interactors"
	"net/http"
)

type TrainingLogHandler struct {
	usecase interactors.ITrainingLogUseCase
}

func NewTrainingLogHandler(lu interactors.ITrainingLogUseCase) *TrainingLogHandler {
	return &TrainingLogHandler{
		usecase: lu,
	}
}

func (lh *TrainingLogHandler) GetLogByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userId := c.Param("userid")

		menus, err := lh.usecase.SelectTrainingLogById(ctx, userId)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, menus)
	}
}
