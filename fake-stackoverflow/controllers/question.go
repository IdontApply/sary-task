package controllers

import (
	"fake-so/httputil"
	m "fake-so/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm/clause"
)

// CreateUser godoc
// @Description
// @Tags         questions
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.Questions
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /question [post]
func CreateQuestion(c *gin.Context) {
	var input m.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	question := m.Question{
		Body:      input.Body,
		UserRefer: input.UserRefer,
		Tital:     input.Tital,
		Tags:      input.Tags,
	}

	if err := m.DB.Omit(clause.Associations).Create(&question).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}
}

// CreateUser godoc
// @Description
// @Tags         questions
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.Questions
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /question [get]
func GetQuestionByID(c *gin.Context) {
	var question m.Question
	var answers []m.Answer
	if err := m.DB.Find(&question).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}
	if err := m.DB.First(&question, c.Param("id")).Error; err != nil {
		log.Println("error: ", err)
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	} else {
		m.DB.Where("question <> ?", c.Param("id")).Find(&answers)
	}
}

// CreateUser godoc
// @Description
// @Tags         questions
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.Questions
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /questions [get]
func GetQuestions(c *gin.Context) {
	var questions []m.Question
	if err := m.DB.Find(&questions).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": questions})
	}

}
