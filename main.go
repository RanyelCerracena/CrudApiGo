package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routeHearth(c *gin.Context) {
	c.JSON(http.StatusOK, )
}

func main() {
	service := gin.Default()
	getRoutes(service)
	service.Run()
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHearth)
	return c
}
