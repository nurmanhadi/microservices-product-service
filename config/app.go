package config

import (
	"product-service/delivery/rest/handler"
	"product-service/delivery/rest/routes"
	"product-service/internal/repository"
	"product-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DependenciesConfig struct {
	DB         *sqlx.DB
	Logger     *logrus.Logger
	Validation *validator.Validate
	Router     *gin.Engine
}

func Setup(deps *DependenciesConfig) {
	// repository
	productRepo := repository.NewProductRepository(deps.DB)

	// service
	productServ := service.NewProductService(deps.Logger, deps.Validation, productRepo)

	// handler
	productHand := handler.NewProductHandler(productServ)

	// routes
	route := &routes.RouteConfig{
		Router:         deps.Router,
		ProductHandler: productHand,
	}
	route.Setup()
}
