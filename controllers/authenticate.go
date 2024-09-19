package controllers

import (
	"net/http"

	"github.com/jaig10/go-crud/models"

	"github.com/gin-gonic/gin"

   "github.com/golang-jwt/jwt/v4"
   "time"
)

var jwtKey = []byte("secretkey")


func Register(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	if err := user.HashPassword(user.Password); err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not Hash Password"})
        return
	}
	newUser := models.User{Username: user.Username, Password: user.Password}

	models.DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"data": "user added successfully"})
}

//Login
func Login(c *gin.Context){
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	if err:= models.DB.Where("username=?", input.Username).First(&user).Error ; err!= nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"User not found"})
		return
	}
	if err:= user.CheckPassword(input.Password); err!=nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid password"})
		return
	}
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(30 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Error in making token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":tokenString})
}

func AllUsers (c *gin.Context){
	var users []models.User
    models.DB.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}
