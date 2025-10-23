package routes

import (
	"product-service/src/rest/handler"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Router          *gin.Engine
	ProductHandler  handler.ProductHandler
	CategoryHandler handler.CategoryHandler
}

func (r *RouteConfig) Setup() {
	api := r.Router.Group("/api")

	product := api.Group("/products")
	product.POST("/", r.ProductHandler.AddProduct)
	product.GET("/", r.ProductHandler.GetAllProducts)
	product.GET("/:id", r.ProductHandler.GetProductByID)
	product.PUT("/:id", r.ProductHandler.UpdateProduct)

	cate := api.Group("/categories")
	cate.POST("/", r.CategoryHandler.AddCategory)
	cate.GET("/", r.CategoryHandler.GetAllCategory)
	cate.POST("/sub", r.CategoryHandler.AddSubCategory)
	cate.PUT("/:id", r.CategoryHandler.UpdateCategoryByID)
	cate.GET("/:id", r.CategoryHandler.GetCategoryByID)
	cate.DELETE("/:id", r.CategoryHandler.DeleteCategoryByID)
}
