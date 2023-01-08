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

type MusclePartOutputFactory func(*gin.Context) ports.MusclePartOutputPort
type MusclePartInputFactory func(ports.MusclePartOutputPort, repository.IMusclePartRepository) ports.MusclePartInputPort
type MusclePartRepositoryFactory func(*sql.DB) repository.IMusclePartRepository

type MusclePartController struct {
	OutputFactory     MusclePartOutputFactory
	InputFactory      MusclePartInputFactory
	RepositoryFactory MusclePartRepositoryFactory
	ClientFactory     database.PostgreSQLConnector
}

func NewMusclePartController(outputFactory MusclePartOutputFactory, inputFactory MusclePartInputFactory, repositoryFactory MusclePartRepositoryFactory, clientFactory database.PostgreSQLConnector) *MusclePartController {
	return &MusclePartController{
		OutputFactory:     outputFactory,
		InputFactory:      inputFactory,
		RepositoryFactory: repositoryFactory,
		ClientFactory:     clientFactory,
	}
}

func (m *MusclePartController) GetAllMuscleParts(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := m.newInputPort(c).GetAllMusclePart(ctx)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (m *MusclePartController) newInputPort(c *gin.Context) ports.MusclePartInputPort {
	outputPort := m.OutputFactory(c)
	repo := m.RepositoryFactory(m.ClientFactory.Conn)
	return m.InputFactory(outputPort, repo)
}
