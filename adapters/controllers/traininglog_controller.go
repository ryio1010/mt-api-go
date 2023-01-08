package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/database"
	"mt-api-go/domain/repository"
	"mt-api-go/usecase/ports"
)

type TrainingLogOutputFactory func(*gin.Context) ports.TrainingLogOutputPort
type TrainingLogInputFactory func(ports.TrainingLogOutputPort, repository.ITrainingLogRepository) ports.TrainingLogInputPort
type TrainingLogRepositoryFactory func(*sql.DB) repository.ITrainingLogRepository

type TrainingLogController struct {
	OutputFactory     TrainingLogOutputFactory
	InputFactory      TrainingLogInputFactory
	RepositoryFactory TrainingLogRepositoryFactory
	ClientFactory     database.PostgreSQLConnector
}

func NewTrainingLogController(outputFactory TrainingLogOutputFactory, inputFactory TrainingLogInputFactory, repositoryFactory TrainingLogRepositoryFactory, clientFactory database.PostgreSQLConnector) *TrainingLogController {
	return &TrainingLogController{
		OutputFactory:     outputFactory,
		InputFactory:      inputFactory,
		RepositoryFactory: repositoryFactory,
		ClientFactory:     clientFactory,
	}
}

func (l *TrainingLogController) GetLogByUserId(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")

		err := l.newInputPort(c).SelectTrainingLogById(ctx, userId)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (l *TrainingLogController) newInputPort(c *gin.Context) ports.TrainingLogInputPort {
	outputPort := l.OutputFactory(c)
	repo := l.RepositoryFactory(l.ClientFactory.Conn)
	return l.InputFactory(outputPort, repo)
}
