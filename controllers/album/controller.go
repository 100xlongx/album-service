package album

import (
	"net/http"
	"web-service-gin/models"
	"web-service-gin/services"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services.GetAllAlbums())
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	album, err := services.GetAlbumByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func Create(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdAlbum := services.CreateAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, createdAlbum)
}
