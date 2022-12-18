package http

import (
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

func InitRouter() *gin.Engine {
	g := gin.Default()

	// DI
	dbConn := infrastructure.NewPostgreSQLConnector()
	repo := postgres.NewRoomRepository(dbConn.Conn)
	svc := service.NewUserService(repo)
	uc := usecase.NewUserUseCase(svc)

	userGroup := g.Group(userApiRoot)
	{
		handler := NewUserHandler(uc)
		// POST LoginAPI
		userGroup.POST("/login", handler.LoginUser())
		// POST AddNewUserAPI
		userGroup.POST("", handler.InsertNewUser())
		// PUT UpdateUserInfoAPI
		userGroup.PUT("", handler.UpdateUser())
	}

	return g
}
