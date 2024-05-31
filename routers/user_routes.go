package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulhakimm/web-go/config"
	"github.com/nulhakimm/web-go/controller"
	"github.com/nulhakimm/web-go/repository"
)

func SetupUserRoutes(app *fiber.App) {
	// Home route
	db, err := config.NewDatabase()
	if err != nil {
		panic(err)
	}
	projectRepo := repository.NewProjectRepo(db)
	projectController := controller.NewProjectController(projectRepo)
	app.Get("/", projectController.RenderHome)

}
