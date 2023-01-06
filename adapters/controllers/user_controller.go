package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/database"
	"mt-api-go/domain/repository"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
)

type IUserController interface {
	LoginUser(ctx context.Context) func(c *gin.Context) error
	InsertNewUser(ctx context.Context) func(c *gin.Context) error
	UpdateUser(ctx context.Context) func(c *gin.Context) error
}

type OutputFactory func(*gin.Context) ports.UserOutputPort
type InputFactory func(ports.UserOutputPort, service.IUserService) ports.UserInputPort
type ServiceFactory func(repository.IUserRepository) service.IUserService
type RepositoryFactory func(*sql.DB) repository.IUserRepository

type UserController struct {
	OutputFactory     OutputFactory
	InputFactory      InputFactory
	ServiceFactory    ServiceFactory
	RepositoryFactory RepositoryFactory
	ClientFactory     database.PostgreSQLConnector
}

func NewUserController(outputFactory OutputFactory, inputFactory InputFactory, serviceFactory ServiceFactory, repositoryFactory RepositoryFactory, clientFactory database.PostgreSQLConnector) *UserController {
	return &UserController{
		OutputFactory:     outputFactory,
		InputFactory:      inputFactory,
		ServiceFactory:    serviceFactory,
		RepositoryFactory: repositoryFactory,
		ClientFactory:     clientFactory,
	}
}

func (u *UserController) LoginUser(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userModel model.User
		err := c.ShouldBind(&userModel)
		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}

		err = u.newInputPort(c).LoginUser(ctx, &userModel)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (u *UserController) InsertNewUser(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userModel model.User
		err := c.ShouldBind(&userModel)
		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}

		err = u.newInputPort(c).InsertNewUser(ctx, &userModel)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (u *UserController) UpdateUser(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userModel model.User
		err := c.ShouldBind(&userModel)
		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}
		err = u.newInputPort(c).UpdateUser(ctx, &userModel)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (u *UserController) newInputPort(c *gin.Context) ports.UserInputPort {
	outputPort := u.OutputFactory(c)
	repo := u.RepositoryFactory(u.ClientFactory.Conn)
	svc := u.ServiceFactory(repo)
	return u.InputFactory(outputPort, svc)
}
