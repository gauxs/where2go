package main

import (
	"net/http"

	"where2go/server/geo"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/map", geo.MapHandler)

	router.Run(":8888")
}
