package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"techsoc-scrumboard-backend/db"
	"techsoc-scrumboard-backend/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db := db.GetDB()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	utils.HandleError(err)
	stmt, err := db.Db.Prepare(fmt.Sprintf("INSERT INTO users (username,password) VALUES ('%s','%s');", credentials.Username, string(hashedPassword)))
	utils.HandleError(err)

	result, err := stmt.Exec()
	fmt.Println(fmt.Sprintf("result from register route: %s", result))
	utils.HandleError(err)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Your account is created successfully"})
}

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db := db.GetDB()
	var password string
	var userId int32
	queryRow := db.Db.QueryRow("SELECT userId,password FROM users WHERE username=?", credentials.Username)
	err := queryRow.Scan(&userId, &password)
	fmt.Println(userId, password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	claims := jwt.MapClaims{
		"username": credentials.Username,
		"userId":   userId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("helloworld"))
	utils.HandleError(err)
	c.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}
