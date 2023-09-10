package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
	return errors.New("Not a existing User")
}

func main() {
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "test",
		})
	})
	server.Run()
}
