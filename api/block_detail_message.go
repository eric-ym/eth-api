package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlockDetailMessageRequest struct {
	Hash  string `json:"hash"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func BlockDetailMessage(c *gin.Context) {
	b := &BlockDetailMessageRequest{}
	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	page := b.Page
	if page <= 0 {
		page = 1
	}

	limit := b.Limit
	if limit <= 0 {
		limit = 20
	}

	data, err := service.GetBlockDetailMessage(b.Hash, page, limit)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)

}
