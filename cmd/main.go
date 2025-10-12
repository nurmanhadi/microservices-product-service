package main

import (
	"product-service/config"
	"product-service/config/database"
)

func main() {
	config.NewEnv()
	logger := config.NewLogger()
	db := database.NewSql()
	validation := config.NewValidator()
	router := config.NewRouter()
	config.Setup(&config.DependenciesConfig{
		DB:         db,
		Logger:     logger,
		Validation: validation,
		Router:     router,
	})

	err := router.Run("0.0.0.0:8081")
	if err != nil {
		logger.Fatalf("failed to start server: %s", err)
	}
}
