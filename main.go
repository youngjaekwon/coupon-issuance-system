package main

import (
	"couponIssuanceSystem/internal/config"
	"couponIssuanceSystem/internal/infra/db"
	"couponIssuanceSystem/routes"
	"log"
)

func main() {
	config.Init()
	db.Init()
	r := routes.SetupRouter()

	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
