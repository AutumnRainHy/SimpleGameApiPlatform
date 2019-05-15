package api

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	server := gin.New()

	game := server.Group("/game")
	{
		game.GET("/GlodFlower", goldFlower) // 炸金花接口；
	}

	return server
}
