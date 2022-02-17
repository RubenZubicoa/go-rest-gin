package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Album struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []Album{
    {Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:3000")
}