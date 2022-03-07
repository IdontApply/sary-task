package controllers

import (
	m "fake-so/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCommentOnQuestion(c *gin.Context) {
	var input m.CommentsOnQuestion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := m.CommentsOnQuestion{
		Body:     input.Body,
		User:     input.User,
		Question: input.Question,
	}

	if err := m.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": comment})
	}
}

func GetCommentsOnQuestion(c *gin.Context) {
	var CommentsOnQuestion []m.CommentsOnQuestion
	if err := m.DB.Find(&CommentsOnQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": CommentsOnQuestion})
	}

}

func CreateCommentOnAnswer(c *gin.Context) {
	var input m.CommentsOnAnswer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := m.CommentsOnAnswer{
		Body:   input.Body,
		User:   input.User,
		Answer: input.Answer,
	}

	if err := m.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": comment})
	}
}

func GetCommentsOnAnswer(c *gin.Context) {
	var CommentsOnAnswer []m.CommentsOnAnswer
	if err := m.DB.Find(&CommentsOnAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": CommentsOnAnswer})
	}

}
