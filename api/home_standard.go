package api

import (
	"eth/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeStandard(c *gin.Context) {
	data, err := service.HomeStandards()

	if err != nil {
		c.JSON(http.StatusOK, &BadReturn{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
