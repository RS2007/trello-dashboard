package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"techsoc-scrumboard-backend/db"
	"techsoc-scrumboard-backend/utils"
)

type Board struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func GetAllBoards(c *gin.Context) {
	Db := db.GetDB().Db
	id, err := strconv.Atoi(c.Param("id"))
	utils.HandleError(err)
	rows, err := Db.Query(fmt.Sprintf("SELECT * FROM boards WHERE workspace = %d", id))
	utils.HandleError(err)

	boardArray := make([]db.BoardStruct, 0, 10)
	for rows.Next() {

		var board db.BoardStruct
		error := rows.Scan(&board.BoardId, &board.Title, &board.Description, &board.Workspace)
		if error != nil {
			utils.HandleError(error)
		}
		boardArray = append(boardArray, board)
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"boards": boardArray})

}

func AddBoard(c *gin.Context) {
	var board Board
	if err := c.ShouldBindJSON(&board); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	Db := db.GetDB()
	id, err := strconv.Atoi(c.Param("id"))
	utils.HandleError(err)
	stmt, err := Db.Db.Prepare(fmt.Sprintf("INSERT INTO boards (title,description,workspace) VALUES ('%s','%s','%d')", board.Title, board.Description, id))
	utils.HandleError(err)
	result, err := stmt.Exec()
	newId, err := result.LastInsertId()
	utils.HandleError(err)
	newBoard := db.BoardStruct{
		Title:       board.Title,
		Description: board.Description,
		BoardId:     int32(newId),
		Workspace:   int32(id),
	}
	utils.HandleError(err)
	c.JSON(http.StatusAccepted, newBoard)
}
