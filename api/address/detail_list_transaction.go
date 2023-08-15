package address

import (
	"eth/api"
	"eth/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddressDetailListTransaction(c *gin.Context) {
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

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	fmt.Println("page: ", page, " limit: ", limit)

	data, err := service.GetAddressTrans(b.Hash, page, limit)

	if err != nil {
		c.JSON(http.StatusOK, &api.BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
