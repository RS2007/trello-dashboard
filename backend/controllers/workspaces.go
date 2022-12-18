package controllers

import (
	"fmt"
	"net/http"
	"techsoc-scrumboard-backend/db"
	"techsoc-scrumboard-backend/utils"

	"github.com/gin-gonic/gin"
)

type Workspace struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func GetAllWorkspaces(c *gin.Context) {
	Db := db.GetDB().Db
	user := c.MustGet("user")
	rows, err := Db.Query(fmt.Sprintf("SELECT * FROM workspaces WHERE user = %d", user))
	utils.HandleError(err)

	workspaceArray := make([]db.WorkspaceStruct, 0, 10)
	for rows.Next() {

		var workspace db.WorkspaceStruct
		error := rows.Scan(&workspace.WorkspaceId, &workspace.Title, &workspace.Description, &workspace.User)
		if error != nil {
			utils.HandleError(error)
		}
		workspaceArray = append(workspaceArray, workspace)
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"workspaces": workspaceArray})

}

func AddWorkspace(c *gin.Context) {
	var workspace Workspace
	if err := c.ShouldBindJSON(&workspace); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	Db := db.GetDB()
	user := c.MustGet("user")
	fmt.Println(user)
	stmt, err := Db.Db.Prepare(fmt.Sprintf("INSERT INTO workspaces (title,description,user) VALUES ('%s','%s','%d')", workspace.Title, workspace.Description, user))
	utils.HandleError(err)
	result, err := stmt.Exec()
	newId, err := result.LastInsertId()
	utils.HandleError(err)
	newWorkspace := db.WorkspaceStruct{
		Title:       workspace.Title,
		Description: workspace.Description,
		WorkspaceId: int32(newId),
		User:        int32(user.(int)),
	}
	utils.HandleError(err)
	c.JSON(http.StatusAccepted, newWorkspace)

}
