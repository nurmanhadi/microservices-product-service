package main

import (
	docs "product-service/docs"
	"product-service/pkg/env"
	"product-service/src/config"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Product Service API
// @version 1.0
// @description This is a product service API server
// @termsOfService http://swagger.io/terms/
// @BasePath /api/products
// @schemes http https
func main() {
	env.NewEnv()
	logger := config.NewLogger()
	db := config.NewSql()
	defer db.Close()
	validation := config.NewValidator()
	router := config.NewRouter()
	conn, ch := config.NewBroker()
	defer conn.Close()
	defer ch.Close()
	config.Setup(&config.DependenciesConfig{
		Ch:         ch,
		DB:         db,
		Logger:     logger,
		Validation: validation,
		Router:     router,
	})

	docs.SwaggerInfo.BasePath = "/api/products"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run(":4001")
	if err != nil {
		logger.Fatalf("failed to start server: %s", err)
	}
}
