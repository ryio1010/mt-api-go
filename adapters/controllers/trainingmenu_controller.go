package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase/interactors"
	"mt-api-go/usecase/model"
	"net/http"
	"strconv"
)

type TrainingMenuHandler struct {
	usecase interactors.ITrainingMenuUseCase
}

func NewTrainingMenuHandler(mu interactors.ITrainingMenuUseCase) *TrainingMenuHandler {
	return &TrainingMenuHandler{
		usecase: mu,
	}
}

func (mh *TrainingMenuHandler) GetMenuById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userId := c.Param("userid")

		menus, err := mh.usecase.GetMenuByUserId(ctx, userId)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, menus)
	}
}

func (mh *TrainingMenuHandler) InsertMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var trainingMenuRequestModel model.TrainingMenuRequest
		err := c.ShouldBind(&trainingMenuRequestModel)
		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}
		fmt.Println(trainingMenuRequestModel)
		insertedMenu, err := mh.usecase.InsertMenu(ctx, &trainingMenuRequestModel)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, insertedMenu)
	}
}

func (mh *TrainingMenuHandler) DeleteMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		menuId := c.Param("menuid")
		idInt, _ := strconv.Atoi(menuId)

		deletedId, err := mh.usecase.DeleteMenu(ctx, idInt)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, deletedId)
	}
}
