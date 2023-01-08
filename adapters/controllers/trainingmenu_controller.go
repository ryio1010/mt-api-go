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
	"strconv"
)

type TrainingMenuOutputFactory func(*gin.Context) ports.TrainingMenuOutputPort
type TrainingMenuInputFactory func(ports.TrainingMenuOutputPort, repository.ITrainingMenuRepository) ports.TrainingMenuInputPort
type TrainingMenuRepositoryFactory func(*sql.DB) repository.ITrainingMenuRepository

type TrainingMenuController struct {
	OutputFactory     TrainingMenuOutputFactory
	InputFactory      TrainingMenuInputFactory
	RepositoryFactory TrainingMenuRepositoryFactory
	ClientFactory     database.PostgreSQLConnector
}

func NewTrainingMenuController(outputFactory TrainingMenuOutputFactory, inputFactory TrainingMenuInputFactory, repositoryFactory TrainingMenuRepositoryFactory, clientFactory database.PostgreSQLConnector) *TrainingMenuController {
	return &TrainingMenuController{
		OutputFactory:     outputFactory,
		InputFactory:      inputFactory,
		RepositoryFactory: repositoryFactory,
		ClientFactory:     clientFactory,
	}
}

func (m *TrainingMenuController) GetMenuById(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")

		err := m.newInputPort(c).GetMenuByUserId(ctx, userId)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (m *TrainingMenuController) InsertMenu(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var trainingMenuRequestModel model.TrainingMenuRequest
		err := c.ShouldBind(&trainingMenuRequestModel)

		if err != nil {
			fmt.Println("バインドエラー")
			fmt.Println(err)
		}

		fmt.Println(trainingMenuRequestModel)
		err = m.newInputPort(c).InsertMenu(ctx, &trainingMenuRequestModel)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (m *TrainingMenuController) DeleteMenu(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		menuId := c.Param("menuid")
		idInt, _ := strconv.Atoi(menuId)

		err := m.newInputPort(c).DeleteMenu(ctx, idInt)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (m *TrainingMenuController) newInputPort(c *gin.Context) ports.TrainingMenuInputPort {
	outputPort := m.OutputFactory(c)
	repo := m.RepositoryFactory(m.ClientFactory.Conn)
	return m.InputFactory(outputPort, repo)
}
