package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func Address(c *gin.Context) {
	b := &AddressListRequest{}

	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, &BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	page := b.Page
	limit := b.Limit

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	data, err := service.GetAddressList(page, limit)

	if err != nil {
		c.JSON(http.StatusOK, &BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)

}
