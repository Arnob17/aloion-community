package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // includes Logger & Recovery middleware

	// Simple route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Aloion Backend!",
		})
	})

	type user struct {
		Name string `json:"name"`
	}

	u := user{Name: "Aloion"}

	// Example route group
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, u)
		})
	}

	// Start server
	r.Run(":8080") // default listens on 0.0.0.0:8080
}
