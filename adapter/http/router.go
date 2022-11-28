package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt-api-go/domain/service"
	"mt-api-go/infrastructure"
	"mt-api-go/infrastructure/postgres"
	"mt-api-go/usecase"
)

const (
	apiVersion  = "/v1"
	userApiRoot = apiVersion + "/user"
	userIdParam = "userid"
)

func InitUserRouter() *gin.Engine {
	g := gin.Default()

	// DI
	dbConn := infrastructure.NewPostgreSQLConnector()
	repo := postgres.NewRoomRepository(dbConn.Conn)
	svc := service.NewUserService(repo)
	uc := usecase.NewUserUseCase(svc)

	userGroup := g.Group(userApiRoot)
	{
		handler := NewUserHandler(uc)
		relativePath := fmt.Sprintf("/:%s", userIdParam)
		userGroup.GET(relativePath, handler.FindUserById())
	}

	return g
}
