package main

import (
	api "GolangTestAPI/src"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(api.AuthMiddleware())

	router.POST("/register", api.Register)
	router.GET("/profile", api.ViewProfile)
	router.PUT("/profile", api.EditProfile)
	router.DELETE("/profile", api.DeleteProfile)
	router.POST("/change-password", api.ChangePassword)

	// Run the server
	router.Run(":8080")
}
