package short

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MakeReq struct{
	Origin string `json:"origin" bind:"required"`
}

func Make(c *gin.Context){
	var req MakeReq
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
	}

	
}

type MatchReq struct{
	Match string `json:"match" bind:"required"`
}

func Match(c *gin.Context){
	var req MatchReq
	err := c.ShouldBindUri(&req)
	if err != nil{
		c.JSON(http.StatusNotFound, nil)
	}

}