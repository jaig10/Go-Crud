package main

import (
	"github.com/jaig10/go-crud/models"
	"github.com/jaig10/go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	models.ConnectDatabase()

	// POST ROUTES
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.FindPosts)
    router.GET("/posts/:id", controllers.FindPost) 
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	// USER ROUTES
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/all-users", controllers.AllUsers)


	router.Run("localhost:8080")
}