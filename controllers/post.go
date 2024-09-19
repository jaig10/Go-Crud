package controllers

import (
	"net/http"
    "github.com/jaig10/go-crud/models"

    "github.com/gin-gonic/gin"
)

// CreatePost Struct
type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdatePostStruct
type UpdatePostInput struct {
    Title   string `json:"title"`
    Content string `json:"content"`
	Author string `json:"author" binding:"required"`
}




// Methods



// CreatePost Method
func CreatePost (c *gin.Context){
	 
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	post := models.Post{Title: input.Title, Content: input.Content, Author: input.Author}
    models.DB.Create(&post)

    c.JSON(http.StatusOK, gin.H{"data": post})

}

//Find All Posts
func FindPosts(c *gin.Context) {
    var posts []models.Post
    models.DB.Find(&posts)

    c.JSON(http.StatusOK, gin.H{"data": posts})
}


//Find Unique Post
func FindPost(c *gin.Context) {
    var post models.Post

    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": post})
}



//Update Post
func UpdatePost(c *gin.Context) {
    var post models.Post
    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    var input UpdatePostInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedPost := models.Post{Title: input.Title, Content: input.Content, Author: input.Author}

    models.DB.Model(&post).Updates(&updatedPost)
    c.JSON(http.StatusOK, gin.H{"data": post})
}

//Delete Post
func DeletePost(c *gin.Context) {
	var post models.Post
	
	if err:= models.DB.Where("id=?", c.Param("id")).First(&post).Error ;err!= nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error":"record not found"})
	}
	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data":"success"})
}