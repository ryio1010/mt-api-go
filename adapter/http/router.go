package http

import (
	"github.com/gin-gonic/gin"
	"mt-api-go/adapter/http/handler"
	"mt-api-go/domain/service"
	"mt-api-go/infrastructure"
	"mt-api-go/infrastructure/postgres"
	"mt-api-go/usecase"
)

const (
	apiVersion  = "/v1"
	userApiRoot = apiVersion + "/user"
	menuApiRoot = apiVersion + "/menu"
	userIdParam = "userid"
)

func InitRouter() *gin.Engine {
	g := gin.Default()

	// DI
	dbConn := infrastructure.NewPostgreSQLConnector()
	userRepository := postgres.NewUserRepository(dbConn.Conn)
	userService := service.NewUserService(userRepository)
	userUseCase := usecase.NewUserUseCase(userService)

	musclePartRepository := postgres.NewMusclePartRepository(dbConn.Conn)
	musclePartService := service.NewMusclePartService(musclePartRepository)
	musclePartUseCase := usecase.NewMusclePartUseCase(musclePartService)

	userGroup := g.Group(userApiRoot)
	{
		userHandler := handler.NewUserHandler(userUseCase)
		// POST LoginAPI
		userGroup.POST("/login", userHandler.LoginUser())
		// POST AddNewUserAPI
		userGroup.POST("", userHandler.InsertNewUser())
		// PUT UpdateUserInfoAPI
		userGroup.PUT("", userHandler.UpdateUser())
	}

	menuGroup := g.Group(menuApiRoot)
	{
		menuHandler := handler.NewMusclePartHandler(musclePartUseCase)
		// GET GetAllMusclePartAPI
		menuGroup.GET("/parts",menuHandler.GetAllMusclePart())
	}

	return g
}
