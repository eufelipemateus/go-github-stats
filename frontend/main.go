package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"

	"github.com/eufelipemateus/go-github-stats/frontend/configs"
	"github.com/eufelipemateus/go-github-stats/frontend/database"
	"github.com/eufelipemateus/go-github-stats/frontend/handlers"
	"github.com/eufelipemateus/go-github-stats/frontend/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	err := config.Load()
	if err != nil {
		panic("Erro on load config.toml")
	}
	database.OpenConnection()
	//database.GenerateDB()

	engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
		ErrorHandler: handlers.ErrorHandler,
    })

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Use(logger.New())

	app.Listen(config.GetServerPort())
}