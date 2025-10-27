package service

import (
	"database/sql"
	"product-service/pkg/response"
	"product-service/src/dto"
	"product-service/src/internal/entity"
	"product-service/src/internal/repository"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ProductCategoryService interface {
	AddProductCategory(request dto.ProductCategoryAddRequest) error
	DeleteProductCategory(id string) error
}
type productCategoryService struct {
	logger                    *logrus.Logger
	validation                *validator.Validate
	productCategoryRepository repository.ProductCategoryRepository
	productRepository         repository.ProductRepository
	categoryRepository        repository.CategoryRepository
}

func NewproductCategoryService(logger *logrus.Logger, validation *validator.Validate, productCategoryRepository repository.ProductCategoryRepository, productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository) ProductCategoryService {
	return &productCategoryService{
		logger:                    logger,
		validation:                validation,
		productCategoryRepository: productCategoryRepository,
		productRepository:         productRepository,
		categoryRepository:        categoryRepository,
	}
}
func (s *productCategoryService) AddProductCategory(request dto.ProductCategoryAddRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.logger.WithError(err).Warn("validation failed for add product category")
		return err
	}
	totalProduct, err := s.productRepository.CountByID(request.ProductID)
	if err != nil {
		s.logger.WithError(err).Error("failed to count user by id")
		return err
	}
	if totalProduct < 1 {
		s.logger.WithError(err).Warn("product not found")
		return response.Except(404, "product not found")
	}
	category, err := s.categoryRepository.FindByID(request.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.WithError(err).Warn("category not found")
			return response.Except(404, "category not found")
		}
		s.logger.WithError(err).Error("failed to find by id")
		return err
	}
	if category.ParentID == nil {
		s.logger.Warn("category not found")
		return response.Except(404, "category not found")
	}
	ps := entity.ProductCategory{
		ProductID:  request.ProductID,
		CategoryID: request.CategoryID,
	}
	if err := s.productCategoryRepository.Insert(ps); err != nil {
		s.logger.WithError(err).Error("failed to insert product category")
		return err
	}
	return nil
}
func (s *productCategoryService) DeleteProductCategory(id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse string to int")
		return err
	}
	totalCategory, err := s.productCategoryRepository.CountByID(int64(newId))
	if err != nil {
		s.logger.WithError(err).Error("failed to check product category existence")
		return err
	}
	if totalCategory < 1 {
		s.logger.Warn("product category not found")
		return response.Except(404, "product category not found")
	}
	if err := s.productCategoryRepository.Delete(int64(newId)); err != nil {
		s.logger.WithError(err).Error("failed to delete product category")
		return err
	}
	return nil
}
