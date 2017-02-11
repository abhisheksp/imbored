package handler

import "github.com/gin-gonic/gin"

func MusicHandler(c *gin.Context) {
	testParam := c.Param("artist")

	c.JSON(200, gin.H{
		"music":  testParam,
	})
}
