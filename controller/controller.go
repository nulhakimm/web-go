package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nulhakimm/web-go/repository"
)

// Sample project data
// var projects = []model.Project{
// 	{Title: "Project 1", Description: "Description for project 1", UrlGithub: "https://github.com/user/project1", UrlDoc: "https://docs.project1.com"},
// 	{Title: "Project 2", Description: "Description for project 2", UrlGithub: "https://github.com/user/project2", UrlDoc: "https://docs.project2.com"},
// 	{Title: "Project 3", Description: "Description for project 3", UrlGithub: "https://github.com/user/project3", UrlDoc: "https://docs.project3.com"},
// }

type ProjectController interface {
	RenderHome(ctx *fiber.Ctx) error
}

type ProjectControllerImpl struct {
	ProjectRepo repository.ProjectRepo
}

func NewProjectController(projectRepo repository.ProjectRepo) ProjectController {
	return &ProjectControllerImpl{
		ProjectRepo: projectRepo,
	}
}

func (controller *ProjectControllerImpl) RenderHome(c *fiber.Ctx) error {
	projects, err := controller.ProjectRepo.FindAll(c.Context())
	if err != nil {
		return fmt.Errorf("failed get all data : %v", err)
	}

	return c.Render("index", fiber.Map{
		"Projects": projects,
	})
}
