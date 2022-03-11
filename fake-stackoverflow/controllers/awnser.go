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
// @Tags         answers
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.Answers
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /answer [post]
func CreateAnswerOfQuestion(c *gin.Context) {
	var input m.Answer
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	answer := m.Answer{
		Body:          input.Body,
		UserRefer:     input.UserRefer,
		QuestionRefer: input.QuestionRefer,
	}

	if err := m.DB.Omit(clause.Associations).Create(&answer).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": answer})
	}

}

// CreateUser godoc
// @Description
// @Tags         answers
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.Answers
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /answers [get]
func GetAnswersOfQuestion(c *gin.Context) {
	var question m.Question
	var answers []m.Answer
	if err := m.DB.First(&question, c.Param("id")).Error; err != nil {
		log.Println("error: ", err)
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	} else {
		m.DB.Where("question <> ?", c.Param("id")).Find(&answers)
	}
}
