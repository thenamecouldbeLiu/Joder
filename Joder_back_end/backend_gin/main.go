package main

import (
	"errors"
	"joder/config"
	"joder/controller"
	"joder/middleware"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"test": "test",
	}
}

func isUserExist(userName string) bool {
	_, isExist := UserData[userName]
	return isExist
}
func checkPassWord(input string, actual string) error {
	if input == actual {
		return nil
	}
	return errors.New("false psw")
}

func Auth(userName string, password string) error {
	if isUserExist(userName) {
		return checkPassWord(password, UserData[userName])
	}
	return errors.New("not a existing User")
}

func load() {
	config.Init()
}

func main() {
	load()
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
