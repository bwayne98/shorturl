package short

import "github.com/gin-gonic/gin"

func Register(route *gin.RouterGroup){
	route.GET("/make", Make)
	route.GET("/match", Match)
}