package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/core"
)

type TransactionHandler struct{}

func (t *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transactionObject core.TransactionRequest
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

	ResponseOKWithPayload(c, transaction)
}

func (t *TransactionHandler) CreateTransactionWithID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseBadRequestWithMessage(c, "invalid ID")
		return
	}

	var transactionObject core.TransactionRequest
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

	ResponseOKWithPayload(c, transaction)
}

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

func (t *TransactionHandler) GetTransactionIDsByType(c *gin.Context) {
	transactionType := c.Param("type")

	transactionIDs, err := core.GetTransactionByType(c.Request.Context(), transactionType)
	if err != nil {
		ResponseServerErrorWithMessage(c, err.Error())
		return
	}

	ResponseOKWithPayload(c, transactionIDs)
}

func (t *TransactionHandler) GetTransactionCumulativeAmount(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseBadRequestWithMessage(c, "invalid ID")
		return
	}

	sum, err := core.GetCumulativeSumByParentID(c.Request.Context(), id)
	if err != nil {
		ResponseServerErrorWithMessage(c, err.Error())
		return
	}

	ResponseOKWithPayload(c, gin.H{"sum": sum})
}
