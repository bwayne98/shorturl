package short

import (
	"shorturl/m/db/model/shorturl"

	"github.com/gin-gonic/gin"
)

func Register(route *gin.RouterGroup, db shorturl.DBTX){

	controller := New(db)

	route.GET("/make", controller.Make)
	route.GET("/match", controller.Match)
}