package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-demo/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/course/list", controllers.GetCourses)
	route.Post("/course/add", controllers.AddCourse)
}
