package main

import (
	"fmt"
	"web-service-gin/controllers/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/albums", album.Get)
	router.GET("/albums/:id", album.GetByID)
	router.POST("/albums", album.Create)

	port := "8080"

	router.Run(fmt.Sprintf("localhost:%s", port))
}
