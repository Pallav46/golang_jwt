package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/Pallav46/golang_jwt/controllers"
	"github.com/Pallav46/golang_jwt/middleware"
)

func UserRoutes(r *gin.Engine) {
	r.Use(middleware.Authenticate())

	r.GET("users", controllers.GetUsers())
	r.GET("users/:userId", controllers.GetUser())
}
