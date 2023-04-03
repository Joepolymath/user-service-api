package routes

import (
	"github.com/Joepolymath/user-service-api/controllers"
	"github.com/gin-gonic/gin"
)

 func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
 }