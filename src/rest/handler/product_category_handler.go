package handler

import (
	"product-service/pkg/response"
	"product-service/src/dto"
	"product-service/src/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler interface {
	AddProductCategory(ctx *gin.Context)
	DeleteProductCategory(ctx *gin.Context)
}
type productCategoryHandler struct {
	productCategoryService service.ProductCategoryService
}

func NewProductCategoryHandler(productCategoryService service.ProductCategoryService) ProductCategoryHandler {
	return &productCategoryHandler{productCategoryService: productCategoryService}
}

// AddProductCategory godoc
// @Summary Add Product Category
// @Description Create relation many to many for product and category
// @Tags product-categories
// @Accept json
// @Produce json
// @Param request body dto.ProductCategoryAddRequest true "Product category add data"
// @Success 201
// @Failure 400
// @Failure 404
// @Router /product-categories [post]
func (h *productCategoryHandler) AddProductCategory(ctx *gin.Context) {
	request := new(dto.ProductCategoryAddRequest)
	if err := ctx.ShouldBind(request); err != nil {
		ctx.Error(response.Except(400, "failed to parse json"))
		return
	}
	err := h.productCategoryService.AddProductCategory(*request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 201, "OK")
}

// DeleteProductCategory godoc
// @Summary Delete Product Category
// @Description Delete relation many to many for product and category
// @Tags product-categories
// @Produce json
// @Param id path string true "Product category id"
// @Success 200
// @Failure 404
// @Router /product-categories/{id} [delete]
func (h *productCategoryHandler) DeleteProductCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.productCategoryService.DeleteProductCategory(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, "OK")
}
