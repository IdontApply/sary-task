package controllers

import (
	m "fake-so/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateTag(c *gin.Context) {
	var input m.Tags
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := m.Tags{Name: input.Name}
	m.DB.Create(&tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}
