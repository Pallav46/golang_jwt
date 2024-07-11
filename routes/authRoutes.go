package routes

import (
	controllers "github.com/Pallav46/golang_jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("users/signup", controllers.Signup())
	r.POST("users/login", controllers.Login())
}