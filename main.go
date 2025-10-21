package main

import (
	"product-service/pkg/env"
	"product-service/src/config"
)

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
	err := router.Run(":4001")
	if err != nil {
		logger.Fatalf("failed to start server: %s", err)
	}
}
