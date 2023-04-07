package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"projectsync/dto"
	"projectsync/service"
)

type ProjectHandler struct {
	svcProject service.ProjectService
}

func NewProjectHandler(svcProject service.ProjectService) *ProjectHandler {
	return &ProjectHandler{svcProject: svcProject}
}

func (p *ProjectHandler) GetAllProjects(ctx *gin.Context) {
	parse, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	listProjects, err := p.svcProject.ListProjects(parse)

	ctx.JSON(200, listProjects)
}

func (p *ProjectHandler) AddProjects(ctx *gin.Context) {
	var projects []dto.NewProject

	if err := ctx.ShouldBindJSON(&projects); err != nil {
		ctx.Error(err)
		return
	}

	if err := p.svcProject.NewProject(projects); err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (p *ProjectHandler) DeleteProject(ctx *gin.Context) {
	str := ctx.Param("id")
	uid, err := uuid.Parse(str)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := p.svcProject.DeleteProject(uid); err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (p *ProjectHandler) UpdateProject(ctx *gin.Context) {
	var projects []dto.UpdateProject
	if err := ctx.ShouldBindJSON(&projects); err != nil {
		ctx.Error(err)
		return
	}

	if err := p.svcProject.UpdateProject(projects); err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.Status(http.StatusOK)
}
