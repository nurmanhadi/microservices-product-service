package main

import (
	"product-service/config"
	"product-service/config/database"
	"product-service/pkg/env"
)

func main() {
	env.NewEnv()
	logger := config.NewLogger()
	db := database.NewSql()
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

	err := router.Run("0.0.0.0:8081")
	if err != nil {
		logger.Fatalf("failed to start server: %s", err)
	}
}
