package api

import (
	"fmt"
	"shorturl/m/api/member"
	"shorturl/m/api/short"
	"shorturl/m/db/store"
	"shorturl/m/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config util.Config
	Query  store.Querier
	Router *gin.Engine
}

func NewServer(config util.Config, query store.Querier, router *gin.Engine) *Server {
	return &Server{
		Config: config,
		Query:  query,
		Router: router,
	}
}

func (s *Server) Run(port string) {
	s.Router.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func (s *Server) SetupRoute() {

	ApiGroup := s.Router.Group("/api")

	memberGroup := ApiGroup.Group("/member")
	member.Default(memberGroup)

	shortGroup := ApiGroup.Group("/short")
	short.Default(shortGroup, s.Query)
}
