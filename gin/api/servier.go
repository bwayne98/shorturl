package api

import (
	"database/sql"
	"fmt"
	"shorturl/m/api/member"
	"shorturl/m/api/short"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db    *sql.DB
	route *gin.Engine
}

func New(db *sql.DB, route *gin.Engine) *Server {
	return &Server{
		db: db,
		route: route,
	}
}

func (s *Server) Run(port string) {
	s.setupRoute()
	s.route.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func (s *Server) setupRoute(){
	
	ApiGroup := s.route.Group("/api")

	memberGroup := ApiGroup.Group("/member")
	member.Default(memberGroup)

	shortGroup := ApiGroup.Group("/short")
	short.Default(shortGroup, s.db)
}


