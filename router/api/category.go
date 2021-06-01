package apihandler

import (
	"net/http"
	"product-api-demo/common/log"
	Model "product-api-demo/model"
	MyType "product-api-demo/type"

	"github.com/gin-gonic/gin"
)

// GetCategories : 取得分類清單
func GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := c.MustGet("Model").(Model.IModel)
		cats, err := model.GetCategories()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.ErrLog.Warn(err)
			return
		}

		c.JSON(http.StatusOK, MyType.RespCategories{
			Categories: cats,
		})
	}
}
