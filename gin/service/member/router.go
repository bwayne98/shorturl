package member

import "github.com/gin-gonic/gin"

func Register(router *gin.RouterGroup){
	router.GET("/login", Login)
	router.GET("/logout", Logout)
}