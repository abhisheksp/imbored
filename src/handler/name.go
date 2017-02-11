package handler

import "github.com/gin-gonic/gin"

func TestHandler(c *gin.Context) {
	testParam := c.Param("testparam")

	c.JSON(200, gin.H{
		"status": "posted",
		"param":  testParam,
	})
}
