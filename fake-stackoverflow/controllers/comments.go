package controllers

import (
	"fake-so/httputil"
	m "fake-so/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Description
// @Tags         comment
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.CommentsOnQuestion
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /question/comment [post]
func CreateCommentOnQuestion(c *gin.Context) {
	var input m.CommentsOnQuestion
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	comment := m.CommentsOnQuestion{
		Body:     input.Body,
		User:     input.User,
		Question: input.Question,
	}

	if err := m.DB.Create(&comment).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": comment})
	}
}

// CreateUser godoc
// @Description
// @Tags         comments
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.CommentsOnQuestion
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /question/{q_id}/comments [get]
func GetCommentsOnQuestion(c *gin.Context) {
	var CommentsOnQuestion []m.CommentsOnQuestion
	if err := m.DB.Find(&CommentsOnQuestion).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": CommentsOnQuestion})
	}

}

// CreateUser godoc
// @Description
// @Tags         vomments
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.CommentsOnAnswer
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /question/answer/comment [post]
func CreateCommentOnAnswer(c *gin.Context) {
	var input m.CommentsOnAnswer
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	comment := m.CommentsOnAnswer{
		Body:   input.Body,
		User:   input.User,
		Answer: input.Answer,
	}

	if err := m.DB.Create(&comment).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"data": comment})
	}
}

// CreateUser godoc
// @Description
// @Tags         comments
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.CommentsOnAnswer
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /question/{q_id}/answer/{a_id}/comments [get]
func GetCommentsOnAnswer(c *gin.Context) {
	var CommentsOnAnswer []m.CommentsOnAnswer
	if err := m.DB.Find(&CommentsOnAnswer).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": CommentsOnAnswer})
	}

}
