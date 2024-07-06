package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/handlers"
)

func AddRoutes(e *gin.Engine) {
	v1 := e.Group("/api/v1")

	transactionHandler := handlers.TransactionHandler{}

	v1.GET("/transaction-service/transaction/:id", transactionHandler.GetTransaction)
}
