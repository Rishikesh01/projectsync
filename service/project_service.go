package service

import (
	"github.com/google/uuid"
	"projectsync/dto"
	"projectsync/repo"
)

type ProjectService interface {
	NewProject([]dto.Project) error
	UpdateProject([]dto.Project) error
	DeleteProject(uuid uuid.UUID) error
}

type projectService struct {
	projectRepo repo.ProjectRepo
}

func NewProjectService(projectRepo repo.ProjectRepo) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (p projectService) NewProject(projects []dto.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p projectService) UpdateProject(projects []dto.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p projectService) DeleteProject(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
