package main

import (
	"github.com/TonySultan/musicstore/controllers"
	"github.com/TonySultan/musicstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/tracks", controllers.GetAllTracks)
	route.POST("/tracks", controllers.CreateTrack)

	route.Run()
}
