package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/abhisheksp/bored/src/handler"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.POST("/testpath/:testparam", handler.TestHandler)
	router.POST("/music/:artists", handler.MusicHandler)


	router.Run(":" + port)
}
