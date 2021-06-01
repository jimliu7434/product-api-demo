package apihandler

import (
	"net/http"
	"product-api-demo/common/log"
	Model "product-api-demo/model"
	MyType "product-api-demo/type"

	"github.com/gin-gonic/gin"
)

// GetProductsByCatID : 根據 CategoryID 取得 商品 清單
func GetProductsByCatID() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := c.MustGet("Model").(Model.IModel)
		catid := c.Param("catid")
		products, err := model.GetProductsByCatID(catid)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.ErrLog.Warn(err)
			return
		}

		c.JSON(http.StatusOK, MyType.RespProducts{
			Products: products,
		})
	}
}

// GetProduct : 根據 productid 取得 商品明細
func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := c.MustGet("Model").(Model.IModel)
		productid := c.Param("id")
		product, err := model.GetProduct(productid)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.ErrLog.Warn(err)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}
