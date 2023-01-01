package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase"
	"mt-api-go/usecase/model"
	"net/http"
)

type UserHandler struct {
	usecase usecase.IUserUseCase
}

func NewUserHandler(uu usecase.IUserUseCase) *UserHandler {
	return &UserHandler{
		usecase: uu,
	}
}

func (uh *UserHandler) LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var userModel model.User
		err := c.ShouldBind(&userModel)
		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}
		fmt.Println(userModel)
		user, err := uh.usecase.LoginUser(ctx, &userModel)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) InsertNewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var userModel model.User
		err := c.ShouldBind(&userModel)
		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}
		fmt.Println(userModel)
		err = uh.usecase.InsertNewUser(ctx, &userModel)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, true)
	}
}

func (uh *UserHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var userModel model.User
		err := c.ShouldBind(&userModel)
		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}
		fmt.Println(userModel)
		err = uh.usecase.UpdateUser(ctx, &userModel)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, true)
	}
}