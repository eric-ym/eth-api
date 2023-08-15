package address

import (
	"eth/api"
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressMessageListRequest struct {
	Hash  string `json:"hash"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func AddressDetailListMessage(c *gin.Context) {
	b := &AddressMessageListRequest{}
	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, &api.BadReturn{
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

	data, err := service.GetAddressMessage(b.Hash, page, limit)

	if err != nil {
		c.JSON(http.StatusOK, &api.BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)

}
