package short

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Register(route *gin.RouterGroup, db *sql.DB){

	controller := New(db)

	route.GET("/make", controller.Make)
	route.GET("/match", controller.Match)
}