package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/abhisheksp/imbored/src/handler"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.POST("/music/:artists", handler.MusicHandler)
	router.POST("/movies/:movies", handler.MovieHandler)

	router.Run(":" + port)
}
