package main

import (
	"net/http"
	"shorturl/m/service/member"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func main() {

	

	Route = gin.Default()

	ApiGroup := Route.Group("/api")

	memberGroup := ApiGroup.Group("/member")
	member.Register(memberGroup)

	ApiGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "wayne",
			"age":  2,
		})
	})

	Route.Run("0.0.0.0:8000")
}
