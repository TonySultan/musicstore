package controllers

import (
	"net/http"

	"github.com/TonySultan/musicstore/models"
	"github.com/gin-gonic/gin"
)

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func GetAllTracks(context *gin.Context) {
	var tracks []models.Track
	models.ConnectDB().Find(&tracks)

	context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

func CreateTrack(context *gin.Context) {
	var input CreateTrackInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{Artist: input.Artist, Title: input.Title}
	models.ConnectDB().Create(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}
