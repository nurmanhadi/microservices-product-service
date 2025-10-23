package handler

import (
	"product-service/pkg/response"
	"product-service/src/dto"
	"product-service/src/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	AddCategory(ctx *gin.Context)
	AddSubCategory(ctx *gin.Context)
	UpdateCategoryByID(ctx *gin.Context)
	GetAllCategory(ctx *gin.Context)
	GetCategoryByID(ctx *gin.Context)
	DeleteCategoryByID(ctx *gin.Context)
}
type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

// AddCategory godoc
// @Summary Add category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param request body dto.CategoryAddRequest true "Category add data"
// @Success 201
// @Failure 400
// @Router /categories [post]
func (h *categoryHandler) AddCategory(ctx *gin.Context) {
	request := new(dto.CategoryAddRequest)
	if err := ctx.ShouldBind(request); err != nil {
		ctx.Error(response.Except(400, "failed to parse json"))
		return
	}
	err := h.categoryService.AddCategory(*request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 201, "OK")
}

// AddSubCategory godoc
// @Summary Add sub category
// @Description Create a sub category by category id
// @Tags categories
// @Accept json
// @Produce json
// @Param request body dto.CategoryAddSubRequest true "Sub category add data"
// @Success 201
// @Failure 400
// @Failure 404
// @Router /categories/sub [post]
func (h *categoryHandler) AddSubCategory(ctx *gin.Context) {
	request := new(dto.CategoryAddSubRequest)
	if err := ctx.ShouldBind(request); err != nil {
		ctx.Error(response.Except(400, "failed to parse json"))
		return
	}
	err := h.categoryService.AddSubCategory(*request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 201, "OK")
}

// UpdateCategoryByID godoc
// @Summary Update category by id
// @Description update name category by id
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category id"
// @Param request body dto.CategoryUpdateSubRequest true "Category update data"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /categories/{id} [put]
func (h *categoryHandler) UpdateCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	request := new(dto.CategoryUpdateSubRequest)
	if err := ctx.ShouldBind(request); err != nil {
		ctx.Error(response.Except(400, "failed to parse json"))
		return
	}
	err := h.categoryService.UpdateCategoryByID(id, *request)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, "OK")
}

// GetAllCategory godoc
// @Summary Get all category
// @Description Get all category with sub category
// @Tags categories
// @Produce json
// @Success 200
// @Router /categories/ [get]
func (h *categoryHandler) GetAllCategory(ctx *gin.Context) {
	result, err := h.categoryService.GetAllCategory()
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}

// GetCategoryByID godoc
// @Summary Get category by id, this endpoint only for category not sub category
// @Description Get a category with sub category
// @Tags categories
// @Produce json
// @param id path string true "Category id"
// @Success 200
// @Success 404
// @Router /categories/{id} [get]
func (h *categoryHandler) GetCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.categoryService.GetCategoryByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, result)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete category by id without sub category or soft delete
// @Tags categories
// @Produce json
// @param id path string true "Category id"
// @Success 200
// @Success 404
// @Router /categories/{id} [delete]
func (h *categoryHandler) DeleteCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.categoryService.DeleteCategoryByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(ctx, 200, "OK")
}
