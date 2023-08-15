package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlockDetailInfoRequest struct {
	Hash string `json:"hash"`
}

func BlockDetailInfo(c *gin.Context) {
	b := &BlockDetailInfoRequest{}
	err := c.ShouldBind(&b)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data, err := service.GetBlockDetail(b.Hash)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
