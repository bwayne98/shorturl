package main

import (
	"database/sql"
	"log"
	"shorturl/m/api"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgresql://wayne:123456qq@pg:5432/shorturl?sslmode=disable")
	if err != nil {
		log.Fatal("cant connect to db:", err)
	}

	route := gin.Default()

	server := api.New(db, route)

	server.Run("8000")
}
