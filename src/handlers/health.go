package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/config"
)

type HealthHandler struct{}

// HealthCheck route for service monitoring
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	ResponseOKWithPayload(c, gin.H{
		"message": "service is up!",
		"version": config.GetVersion(),
		"env":     config.GetEnvType(),
	})
	return
}
