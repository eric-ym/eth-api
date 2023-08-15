package address

import (
	"eth/api"
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressDetailsRequest struct {
	Hash string `json:"hash"`
}

func AddressDetailInfo(c *gin.Context) {
	b := &AddressDetailsRequest{}

	err := c.ShouldBind(b)
	if err != nil {
		c.JSON(http.StatusOK, &api.BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data, err := service.GetAddressDetailInfo(b.Hash)
	if err != nil {
		c.JSON(http.StatusOK, &api.BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
