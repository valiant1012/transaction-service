package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/handlers"
)

func AddRoutes(e *gin.Engine) {
	v1 := e.Group("/api/v1")

	transactionHandler := handlers.TransactionHandler{}

	v1.POST("/transactionservice/transaction", transactionHandler.CreateTransaction)
	v1.PUT("/transactionservice/transaction/:id", transactionHandler.CreateTransactionWithID)
	v1.GET("/transactionservice/transaction/:id", transactionHandler.GetTransactionByID)
	v1.GET("/transactionservice/types/:type", transactionHandler.GetTransactionIDsByType)
	v1.GET("/transactionservice/sum/:id", transactionHandler.GetTransactionCumulativeAmount)
}
