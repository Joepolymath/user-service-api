package routes

import (
	// "github.com/Joepolymath/user-service-api/controllers"
	"github.com/Joepolymath/user-service-api/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
}