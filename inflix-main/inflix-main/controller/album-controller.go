package controller

import (
	"log"
	"net/http"

	"inflix-main/model"
	"inflix-main/service"

	"github.com/gin-gonic/gin"
)

//GetAlbums - List all the Albums
func GetAlbums(c *gin.Context) {
	var albums []model.Album
	err := service.GetAllAlbums(&albums)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, albums)
		log.Printf("Albums List: %v", albums)
	}
}

//CreateAlbum - Create an album
func CreateAlbum(c *gin.Context) {
	var album model.Album
	// var mapAlbum map[string]interface{}
	if c.BindJSON(&album) == nil {
		log.Println(album.Title)
		log.Println(album.Release)
		log.Println(album.Production)
		log.Println(album.Director)
		log.Println(album.IsPremium)
	} else {
		log.Println("Error in binding")
	}

	err := service.CreateAlbum(&album)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, album)
	}
}

//GetAlbum - Get the ablbu detail for a given album id
func GetAlbum(c *gin.Context) {
	id := c.Params.ByName("id")
	var album model.Album
	err := service.GetAlbum(&album, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, album)
	}
}

// UpdateAlbum - Update an existing Album
func UpdateAlbum(c *gin.Context) {
	var album model.Album
	id := c.Params.ByName("id")
	err := service.GetAlbum(&album, id)
	if err != nil {
		c.JSON(http.StatusNotFound, album)
	}
	c.BindJSON(&album)
	err = service.UpdateAlbum(&album, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, album)
	}
}

//DeleteAlbum - Delete an Album
func DeleteAlbum(c *gin.Context) {
	var album model.Album
	id := c.Params.ByName("id")
	err := service.DeleteAlbum(&album, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
