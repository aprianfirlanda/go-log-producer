package main

import (
	"github.com/aprianfirlanda/go-log-producer/internal/config"
	"github.com/aprianfirlanda/go-log-producer/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitLogrus()

	httpServer := fiber.New()
	routes.Register(httpServer)

	config.Log.Info("Starting http server on port 8080")
	if err := httpServer.Listen(":8080"); err != nil {
		config.Log.Fatalf("Error starting http server: %v", err)
	}
}
