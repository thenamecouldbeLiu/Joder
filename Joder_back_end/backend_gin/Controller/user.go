package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserData(ctx *gin.Context, paload interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": paload,
	})
}
