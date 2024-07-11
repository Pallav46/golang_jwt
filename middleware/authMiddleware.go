package middleware

import (
	"net/http"

	"github.com/Pallav46/golang_jwt/helper"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware to authenticate the user
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No token provided",
			})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.UserType)

		c.Next()
	}
}
