package controllers

import (
	"fake-so/httputil"
	m "fake-so/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm/clause"
)

// CreateUser godoc
// @Description
// @Tags         tags
// @Accept       json
// @Produce      json
// @Success      200  {object}  m.Tags
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /tag [post]
func CreateTag(c *gin.Context) {
	var input m.Tag
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	tag := m.Tag{Name: input.Name}
	m.DB.Omit(clause.Associations).Create(&tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}
