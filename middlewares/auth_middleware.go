package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jaig10/go-crud/controllers"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate (c *gin.Context){
	tokenString := c.GetHeader("Authorization")
	fmt.Println(tokenString)
	if tokenString == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Authorixation token required"})
		c.Abort()
		return
	}
	parts:= strings.Split(tokenString, "")
	if len(parts) == 2 && parts[0] == "Bearer"{
		tokenString = parts[1]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signingmethod")
		}
		return []byte(controllers.JwtKey), nil
	})

	if err!=nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token ", "token" : token})
		return
	}
	c.Next()
}