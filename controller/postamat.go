package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"saaws88/config"
	"saaws88/model"
)

func GetAll(c *gin.Context) {
	var postamats []model.Postamat
	config.DB.Find(&postamats)

	c.JSON(http.StatusOK, gin.H{"data": postamats})
}

func GetByNumber(c *gin.Context) {
	var postamat model.Postamat

	if err := config.DB.Where("number = ?", c.Param("number")).First(&postamat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No postamat with number %v found", c.Param("number"))})
	}

	c.JSON(http.StatusOK, gin.H{"data": postamat})
}
