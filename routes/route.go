package routes

import (
	"mampuio-project/config"
	"mampuio-project/controller"
	repository "mampuio-project/repository"

	"github.com/gin-gonic/gin"
)

func WalletRoute(route *gin.Engine) {

	userRepository := repository.NewUserRepository(config.DB)
	userController := controller.NewUserController(*userRepository)

	route.POST("/withdraw", userController.Withdraw)
	route.POST("/balance/:name", userController.Balance)
}
