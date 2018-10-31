package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	Route := gin.Default()
	//stores.Route.Use()
	Route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"title": "hi gocms!!!"})
	})

	Route.Run()
}
