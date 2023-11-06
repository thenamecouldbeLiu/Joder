package main

import (
	"joder/config"
	_ "joder/config"
	"joder/controller"
	_ "joder/dao"
	_ "joder/logger"
	"joder/middleware"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func main() {
	gin.SetMode(config.Val.Mode)
	server := gin.Default()
	server.Use(middleware.CORS())
	//API group for SSO
	api := server.Group("/api")
	{
		api.POST("ouath/google/getUserData", middleware.Auth(), controller.GetUserData)
		api.GET("ouath/google/url", controller.GoogleAccess)
		api.GET("ouath/google/login", controller.GoogleLogin)
		api.POST("/InsertUser", controller.InsertUser)
		api.GET("/getUserData", middleware.Auth(), controller.GetUserData)
		api.POST("/getUnmatched", middleware.Auth(), controller.GetUnmatchedUser)
	}

	server.Run(":" + config.Val.Port)
	log.Infof("serve port: %v \n", config.Val.Port)

}
