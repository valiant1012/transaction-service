package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/core"
	"github.com/valiant1012/transaction-service/src/utility/logger"
)

type TransactionHandler struct{}

// CreateTransaction creates a new transaction and auto-assigns an ID to it. Returns the transaction object as response.
func (t *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transactionObject core.TransactionRequestBody
	err := c.Bind(&transactionObject)
	if err != nil {
		ResponseBadRequestWithMessage(c, "invalid body")
		return
	}

	err = transactionObject.Validate()
	if err != nil {
		ResponseBadRequestWithMessage(c, err.Error())
		return
	}

	transaction, err := core.CreateTransaction(c.Request.Context(), transactionObject)
	if err != nil {
		ResponseServerErrorWithMessage(c, err.Error())
		return
	}

	logger.Infoln("created a new transaction with ID", transaction.ID)

	ResponseOKWithPayload(c, transaction)
}

// CreateTransactionWithID creates a transaction with a pre-defined ID. Returns the transaction object as response.
// This will generate a new ID if there is any collision with existing transaction IDs
func (t *TransactionHandler) CreateTransactionWithID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseBadRequestWithMessage(c, "invalid ID")
		return
	}

	var transactionObject core.TransactionRequestBody
	err = c.Bind(&transactionObject)
	if err != nil {
		ResponseBadRequestWithMessage(c, "invalid body")
		return
	}
	transactionObject.ID = id

	err = transactionObject.Validate()
	if err != nil {
		ResponseBadRequestWithMessage(c, err.Error())
		return
	}

	transaction, err := core.CreateTransaction(c.Request.Context(), transactionObject)
	if err != nil {
		ResponseServerErrorWithMessage(c, err.Error())
		return
	}

	logger.Infoln("created a new transaction with requested ID:", id, "and allocated ID:", transaction.ID)

	ResponseOKWithPayload(c, transaction)
}

// GetTransactionByID Returns the transaction object with required ID as response.
func (t *TransactionHandler) GetTransactionByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseBadRequestWithMessage(c, "invalid ID")
		return
	}

	transaction, found, err := core.GetTransactionByID(c.Request.Context(), id)
	if err != nil {
		ResponseServerErrorWithMessage(c, err.Error())
		return
	}
	if !found {
		ResponseNotFoundWithMessage(c, "transaction not found")
		return
	}

	ResponseOKWithPayload(c, transaction)
}

// GetTransactionIDsByType Returns all transaction IDs matching the required `type`
func (t *TransactionHandler) GetTransactionIDsByType(c *gin.Context) {
	transactionType := c.Param("type")

	transactionIDs, err := core.GetTransactionByType(c.Request.Context(), transactionType)
	if err != nil {
		ResponseServerErrorWithMessage(c, err.Error())
		return
	}

	ResponseOKWithPayload(c, transactionIDs)
}

// GetTransactionCumulativeAmount returns recursive cumulative sum of all transactions linked to a given transaction ID
func (t *TransactionHandler) GetTransactionCumulativeAmount(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseBadRequestWithMessage(c, "invalid ID")
		return
	}

	sum, err := core.GetCumulativeSumByParentTransactionID(c.Request.Context(), id)
	if err != nil {
		ResponseServerErrorWithMessage(c, err.Error())
		return
	}

	ResponseOKWithPayload(c, gin.H{"sum": sum})
}
