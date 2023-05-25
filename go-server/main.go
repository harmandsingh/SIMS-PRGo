package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harmandsingh/SIMS-PRGo/go-server/pkg/configs"
	"github.com/harmandsingh/SIMS-PRGo/go-server/pkg/routes"
)

func main() {
	// Connect to the database
	configs.ConnectDb()

	// Run the fiber app
	app := fiber.New()

	// Setup Routes
	routes.SetupStudentRoutes(app)

	app.Listen(":4000")
}