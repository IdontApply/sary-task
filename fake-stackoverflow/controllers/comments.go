package controllers

import (
	"fake-so/httputil"
	m "fake-so/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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
	var input m.CommentOnQuestion
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	comment := m.CommentOnQuestion{
		Body:          input.Body,
		UserRefer:     input.UserRefer,
		QuestionRefer: input.QuestionRefer,
	}

	if err := m.DB.Omit(clause.Associations).Create(&comment).Error; err != nil {
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
	var CommentsOnQuestion []m.CommentOnQuestion
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
	var input m.CommentOnAnswer
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	comment := m.CommentOnAnswer{
		Body:        input.Body,
		UserRefer:   input.UserRefer,
		AnswerRefer: input.AnswerRefer,
	}

	if err := m.DB.Omit(clause.Associations).Create(&comment).Error; err != nil {
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
	var CommentsOnAnswer []m.CommentOnAnswer
	if err := m.DB.Find(&CommentsOnAnswer).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Println("error: ", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": CommentsOnAnswer})
	}

}
