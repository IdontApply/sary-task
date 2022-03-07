package controllers

import (
	m "fake-so/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateQuestion(c *gin.Context) {
	var input m.Questions
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	question := m.Questions{
		Body:  input.Body,
		User:  input.User,
		Tital: input.Tital,
		Tags:  input.Tags,
	}

	if err := m.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}
}

func GetQuestionByID(c *gin.Context) {
	var question m.Questions
	var answers []m.Answers
	if err := m.DB.Find(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}
	if err := m.DB.First(&question, c.Param("id")); err != nil {
		log.Println("error: ", err)
		c.Writer.WriteHeader(404)
		return
	} else {
		m.DB.Where("question <> ?", c.Param("id")).Find(&answers)
	}
}

func GetQuestions(c *gin.Context) {
	var questions []m.Questions
	if err := m.DB.Find(&questions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": questions})
	}

}
