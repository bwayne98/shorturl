package short

import (
	"shorturl/m/db/store"

	"github.com/gin-gonic/gin"
)

func Default(route *gin.RouterGroup, query store.Querier) {

	controller := New(query)

	route.POST("/make", controller.Make)
	route.POST("/match", controller.Match)
}
