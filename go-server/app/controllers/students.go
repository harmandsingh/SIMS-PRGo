package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/harmandsingh/SIMS-PRGo/go-server/app/models"
	"github.com/harmandsingh/SIMS-PRGo/go-server/pkg/configs"
)

func GetAllStudents(c *fiber.Ctx) error {
	return c.SendString("All Students!")
}

func GetStudentById(c *fiber.Ctx) error {
	return nil
}

func CreateStudent(c *fiber.Ctx) error {
	db := configs.DB.Db
	student := new(models.Student)

	// Validate the body of student and return error if it occurs
	err := c.BodyParser(student)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "invalid request body",
			"data": err,
		})
	}

	err = db.Create(&student).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "error occured while creating student",
			"data": err,	
		})
	}

	// Return the created student
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "error",
		"message": "student successfully created",
		"data": student,	
	})
}