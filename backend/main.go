package main

import (
	"fmt"
	"techsoc-scrumboard-backend/controllers"
	"techsoc-scrumboard-backend/db"
	"techsoc-scrumboard-backend/utils"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.InitializeDB()
	router.Use(utils.CORSMiddleware())
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	workspaces := router.Group("/workspaces", utils.AuthMiddleware)
	{
		workspaces.GET("/all", controllers.GetAllWorkspaces)
		workspaces.POST("/add", controllers.AddWorkspace)
	}

	boards := router.Group("/boards/:id", utils.AuthMiddleware)
	{
		boards.GET("/all", controllers.GetAllBoards)
		boards.POST("/add", controllers.AddBoard)
	}

	cards := router.Group("/cards/:id", utils.AuthMiddleware)
	{
		cards.GET("/all", controllers.GetAllCards)
		cards.POST("/add", controllers.AddCard)
		cards.PUT("/change", controllers.ChangeCardStatus)
	}
	fmt.Println("Hello World!")
	router.Run("localhost:5000")
}
