package apihandler

import (
	"net/http"
	"product-api-demo/common/log"
	Model "product-api-demo/model"
	MyType "product-api-demo/type"

	"github.com/gin-gonic/gin"
)

// GetOrder : 取得訂單明細
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := c.MustGet("Model").(Model.IModel)
		orderid := c.Param("id")
		order, err := model.GetOrder(orderid)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.ErrLog.Warn(err)
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

// NewOrder : 產生新訂單
func NewOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := c.MustGet("Model").(Model.IModel)
		var req MyType.ReqOrderNew
		err := c.BindJSON(&req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			log.AccLog.Info(err)
			return
		}

		order, err := model.NewOrder(req)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.ErrLog.Warn(err)
			return
		}

		c.JSON(http.StatusOK, order)
	}
}
