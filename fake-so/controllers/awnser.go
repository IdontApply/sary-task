package controllers

import (
	m "fake-so/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateAnswerOfQuestion(c *gin.Context) {
	var input m.Answers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	answer := m.Answers{
		Body:     input.Body,
		User:     input.User,
		Question: input.Question,
	}

	if err := m.DB.Create(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": answer})
	}

}

// func GetAnswersOfQuestion(w http.ResponseWriter, r *http.Request) {
// 	var answers []m.Answers
// 	if err := db.Preload("Question").Find(&answers).Error; err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Response-Code", "06")
// 		w.Header().Set("Response-Desc", "Data Not Found")
// 		w.WriteHeader(404)
// 		w.Write([]byte(`{"message":"data not found"}`))
// 		log.Fatalf("Error GetAnswersOfQuestion: %s", err)
// 	} else {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Response-Code", "00")
// 		w.Header().Set("Response-Desc", "Success")
// 		json.NewEncoder(w).Encode(&answers)
// 	}
// }
