package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"users-api/database"
	"users-api/handlers"
)

func main() {
	database.Connect()

	r := gin.Default()

	// libera o front do Vite (porta 5173)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	users := r.Group("/users")
	{
		users.GET("", handlers.ListUsers)
		users.GET("/:id", handlers.GetUser)
		users.POST("", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	r.Run(":8080")
}
