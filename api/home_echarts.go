package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeEcharts(c *gin.Context) {
	data, err := service.HomeEChat()

	if err != nil {
		c.JSON(http.StatusOK, &BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)

}
