package drivers

import (
	"context"
	"github.com/gin-gonic/gin"
	"mt-api-go/adapters/controllers"
	"mt-api-go/adapters/gateways/rdb"
	"mt-api-go/database"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/interactors"
	"mt-api-go/usecase/presenters"
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
	ctx := context.Background()

	// DI
	dbConn := database.NewPostgreSQLConnector()

	// User
	outputPort := presenters.NewUserOutputPort
	inputPort := interactors.NewUserUseCase
	userRepository := rdb.NewUserRepository
	userService := service.NewUserService
	dbClient := database.PostgreSQLConnector{}
	userController := controllers.NewUserController(outputPort, inputPort, userService, userRepository, dbClient)

	// MusclePart
	musclePartRepository := rdb.NewMusclePartRepository(dbConn.Conn)
	musclePartService := service.NewMusclePartService(musclePartRepository)
	musclePartUseCase := interactors.NewMusclePartUseCase(musclePartService)

	// Menu
	trainingMenuRepository := rdb.NewTrainingMenuRepository(dbConn.Conn)
	trainingMenuService := service.NewTrainingMenuService(trainingMenuRepository)
	trainingMenuUseCase := interactors.NewTrainingMenuUseCase(trainingMenuService)

	// Log
	TrainingLogRepository := rdb.NewTrainingLogRepository(dbConn.Conn)
	TrainingLogService := service.NewTrainingLogService(TrainingLogRepository)
	TrainingLogUseCase := interactors.NewTrainingLogUseCase(TrainingLogService)

	userGroup := g.Group(userApiRoot)
	{
		// POST LoginAPI
		userGroup.POST("/login", userController.LoginUser(ctx))
		// POST AddNewUserAPI
		userGroup.POST("", userController.InsertNewUser(ctx))
		// PUT UpdateUserInfoAPI
		userGroup.PUT("", userController.UpdateUser(ctx))
	}

	menuGroup := g.Group(menuApiRoot)
	{
		musclePartHandler := controllers.NewMusclePartHandler(musclePartUseCase)
		trainingMenuHandler := controllers.NewTrainingMenuHandler(trainingMenuUseCase)
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
		TrainingLogHandler := controllers.NewTrainingLogHandler(TrainingLogUseCase)
		// GET GetAllTrainingLogByUserId
		logGroup.GET("/:userid", TrainingLogHandler.GetLogByUserId())
	}

	return g
}
