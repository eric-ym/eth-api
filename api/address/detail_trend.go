package address

import (
	"eth/api"
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddressDetailTrend(c *gin.Context) {
	b := &AddressDetailsRequest{}

	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, &api.BadReturn{
			Code:    http.StatusBadGateway,
			Message: err.Error(),
		})
		return
	}

	data, err := service.GetAddressTrend(b.Hash)
	if err != nil {
		c.JSON(http.StatusOK, &api.BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
