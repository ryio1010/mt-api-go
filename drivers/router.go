package drivers

import (
	"context"
	"github.com/gin-gonic/gin"
	"mt-api-go/adapters/controllers"
	"mt-api-go/adapters/gateways/rdb"
	"mt-api-go/database"
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

	// DB
	dbClient := database.PostgreSQLConnector{}

	// User
	userOutputPort := presenters.NewUserOutputPort
	userInputPort := interactors.NewUserUseCase
	userRepository := rdb.NewUserRepository

	// MusclePart
	musclePartOutputPort := presenters.NewMusclePartOutputPort
	musclePartInputPort := interactors.NewMusclePartUseCase
	musclePartRepository := rdb.NewMusclePartRepository

	// Menu
	trainingMenuOutputPort := presenters.NewTrainingMenuOutputPort
	trainingMenuInputPort := interactors.NewTrainingMenuUseCase
	trainingMenuRepository := rdb.NewTrainingMenuRepository

	// Log
	trainingLogOutputPort := presenters.NewTrainingLogOutputPort
	trainingLogInputPort := interactors.NewTrainingLogUseCase
	trainingLogRepository := rdb.NewTrainingLogRepository

	userGroup := g.Group(userApiRoot)
	{
		userController := controllers.NewUserController(userOutputPort, userInputPort, userRepository, dbClient)

		// POST LoginAPI
		userGroup.POST("/login", userController.LoginUser(ctx))
		// POST AddNewUserAPI
		userGroup.POST("", userController.InsertNewUser(ctx))
		// PUT UpdateUserInfoAPI
		userGroup.PUT("", userController.UpdateUser(ctx))
	}

	menuGroup := g.Group(menuApiRoot)
	{
		musclePartController := controllers.NewMusclePartController(musclePartOutputPort, musclePartInputPort, musclePartRepository, dbClient)
		trainingMenuController := controllers.NewTrainingMenuController(trainingMenuOutputPort, trainingMenuInputPort, trainingMenuRepository, dbClient)

		// GET GetAllMusclePartAPI
		menuGroup.GET("/parts", musclePartController.GetAllMuscleParts(ctx))
		// GET GetMenusByUserIdAPI
		menuGroup.GET("/:userid", trainingMenuController.GetMenuById(ctx))
		// POST InsertNewMenuAPI
		menuGroup.POST("", trainingMenuController.InsertMenu(ctx))
		// DELETE DeleteMenuAPI
		menuGroup.DELETE("/:menuid", trainingMenuController.DeleteMenu(ctx))
	}

	logGroup := g.Group(logApiRoot)
	{
		TrainingLogHandler := controllers.NewTrainingLogController(trainingLogOutputPort, trainingLogInputPort, trainingLogRepository, dbClient)
		// GET GetAllTrainingLogByUserId
		logGroup.GET("/:userid", TrainingLogHandler.GetLogByUserId(ctx))
	}

	return g
}
