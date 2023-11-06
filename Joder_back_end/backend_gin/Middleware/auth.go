package middleware

import (
	"joder/beans"
	client_jwt "joder/client_jwt"
	"joder/commonapi"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie(client_jwt.Token_name)

		userId, err := client_jwt.ParseToken(token)
		if err != nil {
			commonapi.SysError(c, beans.ResponseWrapper{}, "No valid token")
			c.Abort()
			return
		}
		c.Set("userId", userId.UserId)
		c.Next()
	}
}
