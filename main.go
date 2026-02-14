package main

import (
	"mampuio-project/config"
	repository "mampuio-project/repository"
	"mampuio-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(repository.User{})

	routes.WalletRoute(r)

	r.Run(":8080")
}
