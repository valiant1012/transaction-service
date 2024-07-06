package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/config"
)

type HealthHandler struct{}

func (h *HealthHandler) Health(c *gin.Context) {
	ResponseWithPayloadOK(c, gin.H{
		"message": "service is up!",
		"version": config.GetVersion(),
		"env":     config.GetEnvType(),
	})
	return
}
