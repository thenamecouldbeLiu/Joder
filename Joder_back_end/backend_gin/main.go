package main

import (
	"joder/config"
	"joder/controller"
	"joder/logger"
	"joder/middleware"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func main() {
	config.Init()
	logger.Init()
	gin.SetMode(config.Val.Mode)
	server := gin.Default()
	server.Use(middleware.CORS())
	//API group for SSO
	api := server.Group("/api")
	{
		api.GET("ouath/google/url", controller.GoogleAccess)
		api.GET("ouath/google/login", controller.GoogleLogin)
	}

	server.GET("/getUserData", controller.GetUserData)
	server.Run(":" + config.Val.Port)
	log.Infof("serve port: %v \n", config.Val.Port)

}
