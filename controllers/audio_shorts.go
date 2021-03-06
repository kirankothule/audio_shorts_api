package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/kirankothule/audio_shorts_api/domain/audioshort"
	"github.com/kirankothule/audio_shorts_api/services"
	"github.com/kirankothule/audio_shorts_api/utils/audio_utils"
	"github.com/kirankothule/audio_shorts_api/utils/rest_utils"
)

func Get(c *gin.Context) {
	id := strings.TrimSpace(c.Param("audio_id"))
	if len(id) == 0 {
		restErr := rest_utils.NewBadRequestError("Please provied valid id")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, err := services.AudioService.GetAudio(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(audio_utils.IsPublic(c.Request)))
}

func Create(c *gin.Context) {
	var audio audioshort.AudioShort
	if err := c.ShouldBindJSON(&audio); err != nil {
		restErr := rest_utils.NewBadRequestError("Please provied valid data")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, err := services.AudioService.CreateAudio(audio)
	if err != nil {

		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(audio_utils.IsPublic(c.Request)))
}
