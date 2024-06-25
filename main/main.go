package main

import (
	"github.com/gin-gonic/gin"

	"saaws88/config"
	"saaws88/controller"
)

func main() {

	router := gin.Default()

	config.ConnectToDb()

	router.GET("postamat/get", controller.GetAll)
	router.GET("postamat/get/:number", controller.GetByNumber)
	router.POST("postamat/batchcreate", controller.BatchCreate)

	router.Run()

}
