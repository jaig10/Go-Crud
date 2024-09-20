package controllers

import (
	"net/http"

	"github.com/jaig10/go-crud/models"

	"github.com/gin-gonic/gin"

   "github.com/golang-jwt/jwt/v5"
//    "time"
   "fmt"
)

var JwtKey = []byte("secretkey")


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
	c.JSON(http.StatusOK, gin.H{"data": "user added successfully", "user": newUser })
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
	token := jwt.New(jwt.SigningMethodHS256) 
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Error in making token", "err":err, "token": token, "tokenString":tokenString})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":tokenString})
}

func AllUsers (c *gin.Context){
	var users []models.User
    models.DB.Find(&users)

    c.JSON(http.StatusOK, gin.H{"data": users})
}
