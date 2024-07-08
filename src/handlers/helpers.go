package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOKWithPayload(c *gin.Context, payload any) {
	setResponseAndPayload(c, http.StatusOK, payload)
}

func ResponseBadRequestWithPayload(c *gin.Context, payload any) {
	setResponseAndPayload(c, http.StatusBadRequest, payload)
}

func ResponseBadRequestWithMessage(c *gin.Context, message string) {
	setResponseAndPayload(c, http.StatusBadRequest, gin.H{"message": message})
}

func ResponseUnauthorizedWithPayload(c *gin.Context, payload any) {
	setResponseAndPayload(c, http.StatusUnauthorized, payload)
}

func ResponseForbiddenWithPayload(c *gin.Context, payload any) {
	setResponseAndPayload(c, http.StatusForbidden, payload)
}

func ResponseNotFoundWithMessage(c *gin.Context, message any) {
	setResponseAndPayload(c, http.StatusNotFound, gin.H{"message": message})
}

func ResponseNotFoundWithPayload(c *gin.Context, payload any) {
	setResponseAndPayload(c, http.StatusNotFound, payload)
}

func ResponseServerErrorWithMessage(c *gin.Context, message string) {
	setResponseAndPayload(c, http.StatusInternalServerError, gin.H{"message": message})
}

func ResponseServerErrorWithPayload(c *gin.Context, payload any) {
	setResponseAndPayload(c, http.StatusInternalServerError, payload)
}

func setResponseAndPayload(c *gin.Context, responseCode int, payload any) {
	c.JSON(responseCode, payload)
}
