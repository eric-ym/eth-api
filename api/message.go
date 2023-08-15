package api

import (
	"eth/internal/service"
	"eth/libs/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func Message(c *gin.Context) {
	b := &MessageRequest{}
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

	data, err := service.GetMessageList(page, limit, common.MessageTypeAll)
	if err != nil {
		c.JSON(http.StatusOK, BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)

}
