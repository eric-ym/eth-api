package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MessageDetailTransaction(c *gin.Context) {
	b := &MessageDetailRequest{}
	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data, err := service.GetMessageDetailTransactions(b.Hash)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)

}
