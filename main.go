package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// routes.AuthRoutes(router) 

	router.GET("/api/v1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "Server looks good"})
	})

	router.Run(":" + port)
}