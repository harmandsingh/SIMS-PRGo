package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harmandsingh/SIMS-PRGo/go-server/app/controllers"
)

func SetupStudentRoutes(app *fiber.App){
	// Group Routes
	api := app.Group("/api/v1")

	// Student Routes
	api.Get("/students", controllers.GetAllStudents)
	api.Get("/students/:id", controllers.GetStudentById)
	api.Post("/students", controllers.CreateStudent)
}