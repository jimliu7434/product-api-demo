package router

import (
	"github.com/gin-gonic/gin"

	"product-api-demo/common/log"
	Model "product-api-demo/model"
	API "product-api-demo/router/api"
	Middleware "product-api-demo/router/middleware"
)

// SetupRouter :
func SetupRouter(isDebugMode bool, model Model.IModel) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// log
	r.Use(func(c *gin.Context) {
		c.Set("IsDebugMode", isDebugMode)
		c.Set("Model", model)
	})
	r.Use(Middleware.Logger(log.AccLog))

	rAPI := r.Group("/api/")
	{
		rAPI.GET("/cats", API.GetCategories())
		rAPI.GET("/products/:catid", API.GetProductsByCatID())
		rAPI.GET("/product/:id", API.GetProduct())
		rAPI.GET("/order/:id", API.GetOrder())
		rAPI.POST("/order/new", API.NewOrder())
	}

	return r
}
