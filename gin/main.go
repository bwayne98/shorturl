package main

import (
	"database/sql"
	"log"
	"os/user"
	"shorturl/m/service/member"
	"shorturl/m/service/short"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var Route *gin.Engine
var User user.User

func main() {

	db, err := sql.Open("postgres", "postgresql://wayne:123456qq@pg:5432/shorturl?sslmode=disable")
	if err != nil {
		log.Fatal("cant connect to db:", err)
	}

	Route = gin.Default()
	ApiGroup := Route.Group("/api")

	memberGroup := ApiGroup.Group("/member")
	member.Register(memberGroup)

	shortGroup := ApiGroup.Group("/short")
	short.Register(shortGroup, db)

	Route.Run("0.0.0.0:8000")
}
