package main

import (
	"database/sql"
	"fmt"
	"log"
	"shorturl/m/api"
	"shorturl/m/db/store"
	"shorturl/m/util"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("load config error:", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.PGUsername, config.PGPassword, config.PGHost, config.PGPORT, config.PGDatabase))

	if err != nil {
		log.Fatal("cant connect to db:", err)
	}

	query := store.New(db)

	route := gin.Default()

	server := api.NewServer(config, query, route)
	server.SetupRoute()

	server.Run("8000")
}
