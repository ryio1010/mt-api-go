package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/database"
	"mt-api-go/domain/repository"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
)

type UserOutputFactory func(*gin.Context) ports.UserOutputPort
type UserInputFactory func(ports.UserOutputPort, repository.IUserRepository, repository.IBodyCompRepository) ports.UserInputPort
type UserRepositoryFactory func(*sql.DB) repository.IUserRepository
type BodyCompRepositoryFactory func(*sql.DB) repository.IBodyCompRepository

type UserController struct {
	OutputFactory             UserOutputFactory
	InputFactory              UserInputFactory
	UserRepositoryFactory     UserRepositoryFactory
	BodyCompRepositoryFactory BodyCompRepositoryFactory
	ClientFactory             database.PostgreSQLConnector
}

func NewUserController(outputFactory UserOutputFactory, inputFactory UserInputFactory, userRepositoryFactory UserRepositoryFactory, bodyCompRepositoryFactory BodyCompRepositoryFactory, clientFactory database.PostgreSQLConnector) *UserController {
	return &UserController{
		OutputFactory:             outputFactory,
		InputFactory:              inputFactory,
		UserRepositoryFactory:     userRepositoryFactory,
		BodyCompRepositoryFactory: bodyCompRepositoryFactory,
		ClientFactory:             clientFactory,
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
	userRepo := u.UserRepositoryFactory(u.ClientFactory.Conn)
	bodyCompRepo := u.BodyCompRepositoryFactory(u.ClientFactory.Conn)
	return u.InputFactory(outputPort, userRepo, bodyCompRepo)
}
