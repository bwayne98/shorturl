package short

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Default(route *gin.RouterGroup, db *sql.DB) {

	controller := New(db)

	route.POST("/make", controller.Make)
	route.POST("/match", controller.Match)
}
