package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"exchange" : "123456",
		"token" : "123456",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
