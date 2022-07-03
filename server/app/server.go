package main

import (
	"net/http"
	"where2go/server/geo"

	"where2go/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server.Application = server.Where2Go{
		Logger: logger,
	}
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/map", geo.MapHandler)

	router.Run(":7654")
}
