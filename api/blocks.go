package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlocksRequest struct {
	StartBlock int64 `json:"start"`
	Limit      int   `json:"limit"`
}

func Blocks(c *gin.Context) {
	b := &BlocksRequest{}

	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, &BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data, err := service.GetBlocks(b.StartBlock, b.Limit)
	if err != nil {
		c.JSON(http.StatusOK, &BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
