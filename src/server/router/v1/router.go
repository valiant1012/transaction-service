package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/handlers"
	"github.com/valiant1012/transaction-service/src/server/middlewares"
)

func AddRoutes(e *gin.Engine) {
	v1 := e.Group("/api/v1")

	transactionHandler := handlers.TransactionHandler{}

	// Use JWT Auth for these routes
	v1.Use(middlewares.AuthMiddleware())

	v1.POST("/transactionservice/transaction", transactionHandler.CreateTransaction)
	v1.PUT("/transactionservice/transaction/:id", transactionHandler.CreateTransactionWithID) // Not recommended :)
	v1.GET("/transactionservice/transaction/:id", transactionHandler.GetTransactionByID)
	v1.GET("/transactionservice/types/:type", transactionHandler.GetTransactionIDsByType)
	v1.GET("/transactionservice/sum/:id", transactionHandler.GetTransactionCumulativeAmount)
}
