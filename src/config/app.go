package config

import (
	"product-service/src/internal/repository"
	"product-service/src/internal/service"
	"product-service/src/messaging/consumer"
	"product-service/src/rest/handler"
	"product-service/src/rest/routes"
	searchengine "product-service/src/search-engine"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/typesense/typesense-go/v3/typesense"
)

type DependenciesConfig struct {
	DB           *sqlx.DB
	Logger       *logrus.Logger
	Validation   *validator.Validate
	Router       *gin.Engine
	Ch           *amqp091.Channel
	SearchEngine *typesense.Client
}

func Setup(deps *DependenciesConfig) {
	// search engine
	client := searchengine.NewClientSearchEngine(deps.SearchEngine)

	// repository
	productRepo := repository.NewProductRepository(deps.DB)
	catRepo := repository.NewCategoryRepository(deps.DB)
	productCategoryRepo := repository.NewProductCategoryRepository(deps.DB)

	// service
	productServ := service.NewProductService(deps.Logger, deps.Validation, productRepo, client)
	catServ := service.NewCategoryService(deps.Logger, deps.Validation, catRepo)
	productCategoryServ := service.NewproductCategoryService(deps.Logger, deps.Validation, productCategoryRepo, productRepo, catRepo)

	// handler
	productHand := handler.NewProductHandler(productServ)
	catHand := handler.NewCategoryHandler(catServ)
	productCategoryHand := handler.NewProductCategoryHandler(productCategoryServ)

	// routes
	route := &routes.RouteConfig{
		Router:                 deps.Router,
		ProductHandler:         productHand,
		CategoryHandler:        catHand,
		ProductCategoryHandler: productCategoryHand,
	}
	route.Setup()

	// consumer
	productCons := consumer.NewProductConsumer(deps.Logger, deps.Ch, productServ)
	productCons.QueueProduct()
}
