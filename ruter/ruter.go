package ruter

import (
	"eth/api"
	"eth/api/address"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ruter(r *gin.Engine) {

	apis := r.Group("/v1/api")
	{
		apis.GET("test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"aaa": 111, "bbb": 222})
		})

		apis.POST("home/address", api.HomeAddress)
		apis.POST("home/24hours", api.HomeHours)
		apis.POST("home/standard", api.HomeStandard)
		apis.POST("home/trend_echarts", api.HomeEcharts)

		apis.POST("blocks", api.Blocks) //
		apis.POST("blocks/detail/info", api.BlockDetailInfo)
		apis.POST("blocks/detail/message", api.BlockDetailMessage)

		apis.POST("transactions", api.Transactions)

		apis.POST("address", api.Address)
		apis.POST("address/detail/list/transaction", address.AddressDetailListTransaction)
		apis.POST("address/detail/list/message", address.AddressDetailListMessage)
		apis.POST("address/detail/info", address.AddressDetailInfo)
		apis.POST("address/detail/trend", address.AddressDetailTrend)

		apis.POST("message", api.Message)
		apis.POST("message/detail/detail", api.MessageDetail)
		apis.POST("message/detail/info", api.MessageDetailInfo)
		apis.POST("message/detail/transaction", api.MessageDetailTransaction)

	}
}
