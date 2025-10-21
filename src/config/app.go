package config

import (
	"product-service/src/internal/repository"
	"product-service/src/internal/service"
	"product-service/src/messaging/consumer"
	"product-service/src/rest/handler"
	"product-service/src/rest/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type DependenciesConfig struct {
	DB         *sqlx.DB
	Logger     *logrus.Logger
	Validation *validator.Validate
	Router     *gin.Engine
	Ch         *amqp091.Channel
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

	// consumer
	productCons := consumer.NewProductConsumer(deps.Logger, deps.Ch, productServ)
	productCons.QueueProduct()
}
