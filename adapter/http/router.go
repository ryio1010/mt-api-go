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
	logApiRoot  = apiVersion + "/log"
	userIdParam = "userid"
)

func InitRouter() *gin.Engine {
	g := gin.Default()

	// DI
	dbConn := infrastructure.NewPostgreSQLConnector()

	// User
	userRepository := postgres.NewUserRepository(dbConn.Conn)
	userService := service.NewUserService(userRepository)
	userUseCase := usecase.NewUserUseCase(userService)

	// MusclePart
	musclePartRepository := postgres.NewMusclePartRepository(dbConn.Conn)
	musclePartService := service.NewMusclePartService(musclePartRepository)
	musclePartUseCase := usecase.NewMusclePartUseCase(musclePartService)

	// Menu
	trainingMenuRepository := postgres.NewTrainingMenuRepository(dbConn.Conn)
	trainingMenuService := service.NewTrainingMenuService(trainingMenuRepository)
	trainingMenuUseCase := usecase.NewTrainingMenuUseCase(trainingMenuService)

	// Log
	TrainingLogRepository := postgres.NewTrainingLogRepository(dbConn.Conn)
	TrainingLogService := service.NewTrainingLogService(TrainingLogRepository)
	TrainingLogUseCase := usecase.NewTrainingLogUseCase(TrainingLogService)

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
		musclePartHandler := handler.NewMusclePartHandler(musclePartUseCase)
		trainingMenuHandler := handler.NewTrainingMenuHandler(trainingMenuUseCase)
		// GET GetAllMusclePartAPI
		menuGroup.GET("/parts", musclePartHandler.GetAllMusclePart())
		// GET GetMenusByUserIdAPI
		menuGroup.GET("/:userid", trainingMenuHandler.GetMenuById())
		// POST InsertNewMenuAPI
		menuGroup.POST("", trainingMenuHandler.InsertMenu())
		// DELETE DeleteMenuAPI
		menuGroup.DELETE("/:menuid", trainingMenuHandler.DeleteMenu())
	}

	logGroup := g.Group(logApiRoot)
	{
		TrainingLogHandler := handler.NewTrainingLogHandler(TrainingLogUseCase)
		// GET GetAllTrainingLogByUserId
		logGroup.GET("/:userid", TrainingLogHandler.GetLogByUserId())
	}

	return g
}
