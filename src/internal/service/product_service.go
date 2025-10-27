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

type ProductService interface {
	AddProduct(request dto.ProductAddRequest) error
	GetAllProducts() ([]dto.ProductResponse, error)
	GetProductByID(id string) (*dto.ProductResponse, error)
	UpdateProductBulkQuantityByID(consumer []dto.OrderConsumerResponse) error
	UpdateProduct(id string, request dto.ProductUpdateRequest) error
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
func (s *productService) UpdateProductBulkQuantityByID(consumer []dto.OrderConsumerResponse) error {
	// if err := s.validation.Struct(consumer); err != nil {
	// 	s.logger.WithError(err).Warn("validation failed for update status by id")
	// 	return err
	// }
	ids := make([]int, 0, len(consumer))
	for _, x := range consumer {
		ids = append(ids, int(x.ProductID))
	}
	products, err := s.productRepository.FindAllByBulkID(ids)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.WithError(err).Error("products not found")
			return err
		}
		s.logger.WithError(err).Error("failed to find all by bulk id")
		return err
	}
	quantityProducts := make([]dto.ProductUpdateBulkQuantity, 0, len(products))
	for _, x := range products {
		for _, y := range consumer {
			if x.ID == y.ProductID {
				quantityProducts = append(quantityProducts, dto.ProductUpdateBulkQuantity{
					ID:       x.ID,
					Quantity: (x.Quantity - y.Quantity),
				})
			}
		}
	}
	if err := s.productRepository.UpdateBulkQuantityByID(quantityProducts); err != nil {
		s.logger.WithError(err).Error("failed to update bulk quantity by bulk id")
		return err
	}
	return nil
}
func (s *productService) UpdateProduct(id string, request dto.ProductUpdateRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.logger.WithError(err).Warn("validation failed for add product")
		return err
	}
	newId, err := strconv.Atoi(id)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse string to int")
		return err
	}
	totalProduct, err := s.productRepository.CountByID(int64(newId))
	if err != nil {
		s.logger.WithError(err).Error("failed to count user by id")
		return err
	}
	if totalProduct < 1 {
		s.logger.WithError(err).Warn("product not found")
		return response.Except(404, "product not found")
	}
	if err := s.productRepository.UpdateProduct(int64(newId), request); err != nil {
		s.logger.WithError(err).Error("failed to update product")
		return err
	}
	return nil
}
