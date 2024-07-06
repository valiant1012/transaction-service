package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseWithPayloadOK(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusOK, payload)
}

func ResponseWithPayloadBadRequest(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusBadRequest, payload)
}

func ResponseWithPayloadUnauthorized(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusUnauthorized, payload)
}

func ResponseWithPayloadForbidden(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusForbidden, payload)
}

func ResponseWithPayloadNotFound(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusNotFound, payload)
}

func ResponseWithPayloadServerError(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusInternalServerError, payload)
}
