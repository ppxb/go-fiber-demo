package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-demo/app/models"
	"go-fiber-demo/app/queries"
)

func GetCourses(c *fiber.Ctx) error {
	courses, err := queries.GetCourses()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "courses not found",
			"total":   0,
			"courses": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"total":   len(courses),
		"courses": courses,
	})
}

func AddCourse(c *fiber.Ctx) error {
	var course models.Course

	if err := c.BodyParser(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := queries.AddCourse(&course); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "success",
	})
}
