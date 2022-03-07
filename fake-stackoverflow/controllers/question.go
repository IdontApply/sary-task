package controllers

import (
	"fake-so/httputil"
	m "fake-so/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	var input m.Questions
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	question := m.Questions{
		Body:  input.Body,
		User:  input.User,
		Tital: input.Tital,
		Tags:  input.Tags,
	}

	if err := m.DB.Create(&question).Error; err != nil {
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
	var question m.Questions
	var answers []m.Answers
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
	var questions []m.Questions
	if err := m.DB.Find(&questions).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": questions})
	}

}
