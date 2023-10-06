package services

import (
	"fmt"
	"strconv"
	"web-service-gin/models"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllAlbums() []models.Album {
	return albums
}

func GetAlbumByID(id string) (*models.Album, error) {
	for _, album := range albums {
		if album.ID == id {
			return &album, nil
		}
	}

	return nil, fmt.Errorf("album with ID %s not found", id)
}

func CreateAlbum(a models.Album) *models.Album {
	newID := strconv.Itoa(len(albums) + 1)
	a.ID = newID
	albums = append(albums, a)
	return &a
}
