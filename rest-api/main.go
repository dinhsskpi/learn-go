package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for i := range albums {
		if albums[i].ID == id {
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func createAlbum(c *gin.Context) {
	var newAlbum album

	if error := c.BindJSON(&newAlbum); error != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	server := gin.Default()

	server.GET("albums", getAlbums)
	server.GET("albums/:id", getAlbumsById)
	server.POST("albums", createAlbum)

	server.Run("localhost:8080")
}
