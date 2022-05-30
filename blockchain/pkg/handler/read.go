package handler

import (
	"blockchain/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Chain(c *gin.Context) {
	res := service.Chain()
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}
func Check(c *gin.Context) {
	res := service.CheckIntegrity()
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}
func MakeOutputConfig(c *gin.Context) {
	transaction := service.Transaction{
		User: c.Query("user"),
		VmId: c.Query("vmId"),
	}
	o := service.MakeOutputConfig(transaction)
	if o.Account == "" {
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"data": "error",
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": o,
		})
	}
}
