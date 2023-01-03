package presenters

import (
	"github.com/gin-gonic/gin"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
	"net/http"
)

type UserPresenter struct {
	ctx *gin.Context
}

func NewUserOutputPort(context *gin.Context) ports.UserOutputPort {
	return &UserPresenter{
		ctx: context,
	}
}

func (u *UserPresenter) OutputUser(user *model.User) error {
	u.ctx.JSON(http.StatusOK, user)
	return nil
}

func (u *UserPresenter) OutputError(err error) error {
	u.ctx.JSON(http.StatusInternalServerError, err)
	return err
}
