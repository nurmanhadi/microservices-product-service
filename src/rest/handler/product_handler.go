package handler

import (
	"product-service/pkg/response"
	"product-service/src/dto"
	"product-service/src/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
}
type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

// AddProduct godoc
// @Summary Add Product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.ProductAddRequest true "Product add data"
// @Success 201
// @Failure 400
// @Router / [post]
func (h *productHandler) AddProduct(ctx *gin.Context) {
	request := new(dto.ProductAddRequest)
	if err := ctx.ShouldBind(request); err != nil {
		ctx.Error(response.Except(400, "failed to parse json"))
		return
	}
	err := h.productService.AddProduct(*request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 201, "OK")
}

// GetAllProducts godoc
// @Summary Get all Products
// @Description Get list Products
// @Tags products
// @Produce json
// @Success 200
// @Router / [get]
func (h *productHandler) GetAllProducts(ctx *gin.Context) {
	result, err := h.productService.GetAllProducts()
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}

// GetProductByID godoc
// @Summary Get product by id
// @Description Get product by id
// @Tags products
// @Produce json
// @Param id path string true "Product id"
// @Success 200
// @Router /{id} [get]
func (h *productHandler) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.productService.GetProductByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}
