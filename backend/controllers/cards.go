package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"techsoc-scrumboard-backend/db"
	"techsoc-scrumboard-backend/utils"
)

type Card struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

func GetAllCards(c *gin.Context) {
	Db := db.GetDB().Db
	id, err := strconv.Atoi(c.Param("id"))
	utils.HandleError(err)
	rows, err := Db.Query(fmt.Sprintf("SELECT * FROM cards WHERE board = %d", id))
	utils.HandleError(err)

	// rowsBoard, err := Db.Query(fmt.Sprintf("SELECT title FROM boards WHERE boardId = %d", id))
	// utils.HandleError(err)

	// var title string
	// for rowsBoard.Next() {
	// 	err = rows.Scan(&title)
	// 	if err != nil {
	// 		utils.HandleError(err)
	// 	}
	// }

	cardArray := make([]db.CardStruct, 0, 10)
	for rows.Next() {

		var card db.CardStruct
		error := rows.Scan(&card.CardId, &card.Title, &card.Description, &card.Status, &card.Board)
		if error != nil {
			utils.HandleError(error)
		}
		cardArray = append(cardArray, card)
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"cards": cardArray})

}

func AddCard(c *gin.Context) {
	var card Card
	if err := c.ShouldBindJSON(&card); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	Db := db.GetDB()
	id, err := strconv.Atoi(c.Param("id"))
	utils.HandleError(err)
	stmtDebug := fmt.Sprintf("INSERT INTO cards (title,description,status,board) VALUES ('%s','%s','%s','%d')", card.Title, card.Description, card.Status, id)
	fmt.Println(stmtDebug)
	stmt, err := Db.Db.Prepare(stmtDebug)
	utils.HandleError(err)
	result, err := stmt.Exec()
	fmt.Println(result)
	newId, err := result.LastInsertId()
	utils.HandleError(err)
	newCard := db.CardStruct{
		Title:       card.Title,
		Description: card.Description,
		Status:      card.Status,
		CardId:      int32(newId),
		Board:       int32(id),
	}
	utils.HandleError(err)
	c.JSON(http.StatusAccepted, newCard)
}

func ChangeCardStatus(c *gin.Context) {
	type changeCardStruct struct {
		CardId    int32  `json:"cardId" required:"binding"`
		NewStatus string `json:"newStatus" required:"binding"`
	}
	var card changeCardStruct
	if err := c.ShouldBindJSON(&card); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	Db := db.GetDB().Db
	stmtDebug := fmt.Sprintf("UPDATE cards SET status = '%s' WHERE cardId = %d", card.NewStatus, card.CardId)
	fmt.Println(stmtDebug)

	stmt, err := Db.Prepare(stmtDebug)
	utils.HandleError(err)
	result, err := stmt.Exec()
	fmt.Println(result)
	utils.HandleError(err)
	c.JSON(http.StatusAccepted, gin.H{"message": "status change successful"})

}
