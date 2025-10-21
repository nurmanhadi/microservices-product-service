package routes

import (
	"product-service/src/rest/handler"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Router         *gin.Engine
	ProductHandler handler.ProductHandler
}

func (r *RouteConfig) Setup() {
	api := r.Router.Group("/api")

	product := api.Group("/products")
	product.POST("/", r.ProductHandler.AddProduct)
	product.GET("/", r.ProductHandler.GetAllProducts)
	product.GET("/:id", r.ProductHandler.GetProductByID)
}
