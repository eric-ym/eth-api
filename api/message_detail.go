package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageDetailRequest struct {
	Hash string `json:"hash"`
}

func MessageDetail(c *gin.Context) {
	b := &MessageDetailRequest{}
	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data, err := service.GetMessageDetail(b.Hash)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
