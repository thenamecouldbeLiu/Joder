package commonapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, payload interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"resultBody":    payload,
		"resultCode":    "0000",
		"resultMessage": message,
	})
}

func SysError(c *gin.Context, payload interface{}, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{

		"resultCode":    "9999",
		"resultMessage": message,
	})
}
