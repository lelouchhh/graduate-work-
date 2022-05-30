package handler

import (
	"blockchain/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewBlock(c *gin.Context) {
	transaction := service.Transaction{
		User:    c.Query("user"),
		Action:  c.Query("action"),
		Account: c.Query("account"),
		Ssh:     c.Query("SSH"),
		Ram:     c.Query("ram"),
		Disk:    c.Query("disk"),
		Cpu:     c.Query("cpu"),
		Pass:    c.Query("pass"),
		VmId:    c.Query("vmId"),
		VmIP:    c.Query("IP"),
	}
	_ = service.CreateBlock(transaction)
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": transaction,
	})
}
