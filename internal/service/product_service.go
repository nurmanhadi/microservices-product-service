package service

import (
	"database/sql"
	"product-service/internal/dto"
	"product-service/internal/entity"
	"product-service/internal/repository"
	"product-service/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ProductService interface {
	AddProduct(request dto.ProductAddRequest) error
	GetAllProducts() ([]dto.ProductResponse, error)
	GetProductByID(id string) (*dto.ProductResponse, error)
	UpdateProductQuantity(productID, quantity string) error
}
type productService struct {
	logger            *logrus.Logger
	validation        *validator.Validate
	productRepository repository.ProductRepository
}

func NewProductService(logger *logrus.Logger, validation *validator.Validate, productRepository repository.ProductRepository) ProductService {
	return &productService{
		logger:            logger,
		validation:        validation,
		productRepository: productRepository,
	}
}
func (s *productService) AddProduct(request dto.ProductAddRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.logger.WithError(err).Warn("validation failed for add product")
		return err
	}
	product := &entity.Product{
		Name:     request.Name,
		Quantity: request.Quantity,
		Price:    request.Price,
	}
	if err := s.productRepository.Insert(*product); err != nil {
		s.logger.WithError(err).Error("Failed to insert new product")
		return err
	}
	return nil
}
func (s *productService) GetAllProducts() ([]dto.ProductResponse, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		s.logger.WithError(err).Error("Failed to find all products from repository")
		return nil, err
	}
	resp := make([]dto.ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, dto.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Quantity:  product.Quantity,
			Price:     product.Price,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	return resp, nil
}

func (s *productService) GetProductByID(productID string) (*dto.ProductResponse, error) {
	newId, err := strconv.Atoi(productID)
	if err != nil {
		s.logger.Warnf("failed to parse productID %s to int", productID)
		return nil, err
	}
	product, err := s.productRepository.FindByID(int64(newId))
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Warnf("Product with ID %d not found", newId)
			return nil, response.Except(404, "product not found")
		}
		s.logger.WithError(err).Errorf("Failed to find product by ID %d", newId)
		return nil, err
	}
	resp := &dto.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	return resp, nil
}
func (s *productService) UpdateProductQuantity(productID, quantity string) error {
	newId, err := strconv.Atoi(productID)
	if err != nil {
		s.logger.Warnf("failed to parse productID %s to int", productID)
		return err
	}
	newQuantity, err := strconv.Atoi(productID)
	if err != nil {
		s.logger.Warnf("failed to parse quantity %s to int", quantity)
		return err
	}
	totalProduct, err := s.productRepository.CountByID(int64(newId))
	if err != nil {
		s.logger.WithError(err).Error("failed to count product")
		return err
	}
	if totalProduct < 1 {
		s.logger.WithError(err).Error("product not found")
		return response.Except(404, "product not found")
	}
	if err := s.productRepository.UpdateQuantity(int64(newId), int64(newQuantity)); err != nil {
		s.logger.WithError(err).Error("Failed to update product quantity")
		return err
	}
	return nil
}
