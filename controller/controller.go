package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nulhakimm/web-go/model"
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
	RenderProject(ctx *fiber.Ctx) error
	CreateProject(ctx *fiber.Ctx) error
}

type ProjectControllerImpl struct {
	ProjectRepo repository.ProjectRepo
}

func NewProjectController(projectRepo repository.ProjectRepo) ProjectController {
	return &ProjectControllerImpl{
		ProjectRepo: projectRepo,
	}
}

func (controller *ProjectControllerImpl) RenderHome(ctx *fiber.Ctx) error {
	projects, err := controller.ProjectRepo.FindAll(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed get all data : %v", err)
	}

	return ctx.Render("index", fiber.Map{
		"Projects": projects,
	})
}

func (controller *ProjectControllerImpl) RenderProject(ctx *fiber.Ctx) error {
	return ctx.Render("project_create", nil)
}

func (controller *ProjectControllerImpl) CreateProject(ctx *fiber.Ctx) error {

	project := model.Project{
		Title:       ctx.FormValue("title"),
		UrlGithub:   ctx.FormValue("url_github"),
		UrlDoc:      ctx.FormValue("url_doc"),
		Description: ctx.FormValue("description"),
		Image:       "assets/projects1.jpg",
	}

	err := controller.ProjectRepo.Save(ctx.Context(), &project)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
	})

}
