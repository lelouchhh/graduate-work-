package main

import (
	h "blockchain/pkg/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	write := router.Group("/write")
	{
		write.GET("/new", h.NewBlock)
		//write.POST("/action_refresh", h.ActionUserRefresh)
	}
	read := router.Group("/read")
	{
		read.GET("/chain", h.Chain)
		read.GET("/check", h.Check)
		read.GET("/output", h.MakeOutputConfig)
		//read.POST("get_entity", h.GetEntity)
	}
	router.Run("localhost:8082")
}
