package controller

import (
	"joder/commonapi"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	commonapi.Success(c, gin.H{
		"user": "001",
	})
}
