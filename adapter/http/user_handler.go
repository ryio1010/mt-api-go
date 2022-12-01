package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase"
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

func (uh *UserHandler) FindUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userId := c.Param("userid")

		user, err := uh.usecase.FindUserById(ctx, userId)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, user)
	}
}
