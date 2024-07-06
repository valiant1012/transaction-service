package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/handlers"
	v1 "github.com/valiant1012/transaction-service/src/server/routers/v1"
)

func AddRoutes(e *gin.Engine) {
	healthHandler := handlers.HealthHandler{}

	// Add generic routes
	e.GET("/health", healthHandler.Health)

	// Add v1 routes
	v1.AddRoutes(e)
}
