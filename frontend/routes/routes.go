package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/eufelipemateus/go-github-stats/frontend/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/u/:user", controllers.GetUser)
}