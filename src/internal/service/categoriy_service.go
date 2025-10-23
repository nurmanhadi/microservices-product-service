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

type CategoryService interface {
	AddCategory(req dto.CategoryAddRequest) error
	AddSubCategory(req dto.CategoryAddSubRequest) error
	UpdateCategoryByID(id string, req dto.CategoryUpdateSubRequest) error
	DeleteCategoryByID(id string) error
	GetCategoryByID(id string) (*dto.CategoryResponse, error)
	GetAllCategory() ([]dto.CategoryResponse, error)
}
type categoryService struct {
	logger             *logrus.Logger
	validation         *validator.Validate
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(logger *logrus.Logger, validation *validator.Validate, categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		logger:             logger,
		validation:         validation,
		categoryRepository: categoryRepository,
	}
}
func (s *categoryService) AddCategory(req dto.CategoryAddRequest) error {
	if err := s.validation.Struct(&req); err != nil {
		s.logger.WithError(err).Warn("validation failed for insert category")
		return err
	}
	cat := entity.Category{
		Name: req.Name,
	}
	err := s.categoryRepository.InsertCategory(cat)
	if err != nil {
		s.logger.WithError(err).Error("failed to insert category")
		return err
	}
	return nil
}
func (s *categoryService) AddSubCategory(req dto.CategoryAddSubRequest) error {
	if err := s.validation.Struct(&req); err != nil {
		s.logger.WithError(err).Warn("validation failed for insert sub category")
		return err
	}
	count, err := s.categoryRepository.CountByID(req.CategoryID)
	if err != nil {
		s.logger.WithError(err).Error("failed to count by id")
		return err
	}
	if count == 0 {
		s.logger.Warn("category not found")
		return response.Except(404, "category not found")
	}
	cat := entity.Category{
		Name:     req.Name,
		ParentID: &req.CategoryID,
	}
	err = s.categoryRepository.InsertSubCategory(cat)
	if err != nil {
		s.logger.WithError(err).Error("failed to insert sub category")
		return err
	}
	return nil
}
func (s *categoryService) UpdateCategoryByID(id string, req dto.CategoryUpdateSubRequest) error {
	newID, err := strconv.Atoi(id)
	if err != nil {
		s.logger.WithError(err).Error("failed to convert string to integer")
		return err
	}
	if err := s.validation.Struct(&req); err != nil {
		s.logger.WithError(err).Warn("validation failed to update by id")
		return err
	}

	count, err := s.categoryRepository.CountByID(newID)
	if err != nil {
		s.logger.WithError(err).Error("failed to check category existence")
		return err
	}
	if count == 0 {
		s.logger.Warn("category not found")
		return response.Except(404, "category not found")
	}

	cat := entity.Category{
		ID:   newID,
		Name: req.Name,
	}

	if err := s.categoryRepository.UpdateByID(cat); err != nil {
		s.logger.WithError(err).Error("failed to update category")
		return err
	}
	return nil
}
func (s *categoryService) DeleteCategoryByID(id string) error {
	newID, err := strconv.Atoi(id)
	if err != nil {
		s.logger.WithError(err).Error("failed to convert string to integer")
		return err
	}
	count, err := s.categoryRepository.CountByID(newID)
	if err != nil {
		s.logger.WithError(err).Warn("failed to check category existence")
		return err
	}
	if count == 0 {
		s.logger.Warn("category not found")
		return response.Except(404, "category not found")
	}

	if err := s.categoryRepository.DeleteByID(newID); err != nil {
		s.logger.WithError(err).Error("failed to delete category")
		return err
	}
	return nil
}
func (s *categoryService) GetAllCategory() ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepository.FindAll()
	if err != nil {
		s.logger.WithError(err).Error("failed to find all categories")
		return nil, err
	}
	var cat []dto.CategoryResponse
	for _, x := range categories {
		if x.ParentID == nil {
			var subCat []dto.CategoryResponse
			for _, y := range categories {
				if y.ParentID != nil && *y.ParentID == x.ID {
					subCat = append(subCat, dto.CategoryResponse{
						ID:   y.ID,
						Name: y.Name,
					})
				}
			}
			cat = append(cat, dto.CategoryResponse{
				ID:            x.ID,
				Name:          x.Name,
				SubCategories: subCat,
			})
		}
	}
	return cat, nil
}
func (s *categoryService) GetCategoryByID(id string) (*dto.CategoryResponse, error) {
	newID, err := strconv.Atoi(id)
	if err != nil {
		s.logger.WithError(err).Error("failed to convert string to integer")
		return nil, err
	}
	category, err := s.categoryRepository.FindByID(newID)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.WithError(err).Warn("category not found")
			return nil, response.Except(404, "category not found")
		}
		s.logger.WithError(err).Error("failed to find by id")
		return nil, err
	}
	if category.ParentID != nil {
		s.logger.Warn("category not found")
		return nil, response.Except(404, "category not found")
	}
	subCategory, err := s.categoryRepository.FindAllSubCategoryByID(newID)
	if err != nil {
		if err != sql.ErrNoRows {
			s.logger.WithError(err).Error("failed to find all sub category by id")
			return nil, err
		}
	}
	subCat := make([]dto.CategoryResponse, 0, len(subCategory))
	if len(subCategory) != 0 {
		for _, x := range subCategory {
			subCat = append(subCat, dto.CategoryResponse{
				ID:   x.ID,
				Name: x.Name,
			})
		}
	}
	resp := &dto.CategoryResponse{
		ID:            category.ID,
		Name:          category.Name,
		SubCategories: subCat,
	}
	return resp, nil
}
