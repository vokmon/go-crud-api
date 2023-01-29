package main

import (
	"go-crud-api/controllers"
	"go-crud-api/initializers"
	"go-crud-api/middlewares"

	"github.com/gin-gonic/gin"
)

// This runs before main
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
