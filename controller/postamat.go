package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"saaws88/config"
	"saaws88/model"
	"saaws88/util"
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

func BatchCreate(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Can't get file"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Can't read file"})
		return
	}
	defer f.Close()

	sheet, err := util.ParseCsv(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Can't parse csv file"})
		return
	}

	count := 0
	for _, line := range sheet {
		p := util.BuildFromLine(line)
		config.DB.Create(&p)
		count++
	}

	c.JSON(http.StatusCreated, gin.H{"data": count})

}
