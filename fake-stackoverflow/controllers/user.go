package controllers

import (
	"fake-so/httputil"
	m "fake-so/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateUser godoc
// @Description  Create a User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Users ID"
// @Success      200  {object}  m.Users
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /user [post]
func CreateUser(c *gin.Context) {
	var input m.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	user := m.Users{Name: input.Name}
	m.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
