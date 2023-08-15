package api

import (
	"eth/internal/service"
	"eth/libs/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeAddress(c *gin.Context) {
	data, err := service.GetHomeAddressList(common.HomeLimit)

	if err != nil {
		c.JSON(http.StatusOK, &BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
