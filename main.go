package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-demo/pkg/database"
	"go-fiber-demo/pkg/routes"
	"go-fiber-demo/pkg/utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.InitDb()

	app := fiber.New()

	routes.PublicRoutes(app)

	utils.StartServer(app)
}
