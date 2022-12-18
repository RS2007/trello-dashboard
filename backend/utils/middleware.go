package utils

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeaderValue := c.GetHeader("Authorization")
	token := strings.Split(authHeaderValue, " ")[1]
	if authHeaderValue != "" {
		claims := jwt.MapClaims{}
		user, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("helloworld"), nil
		})
		HandleError(err)
		c.Set("user", int(user.Claims.(jwt.MapClaims)["userId"].(float64)))
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
}
