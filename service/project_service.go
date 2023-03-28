package service

import (
	"github.com/Rishikesh01/gofp"
	"github.com/google/uuid"
	"projectsync/dto"
	"projectsync/model"
	"projectsync/repo"
)

type ProjectService interface {
	NewProject([]dto.Project) error
	UpdateProject([]dto.Project) error
	DeleteProject(uuid uuid.UUID) error
	ListProjects(uuid uuid.UUID) ([]dto.Project, error)
}

type projectService struct {
	projectRepo repo.ProjectRepo
}

func (p *projectService) ListProjects(uuid2 uuid.UUID) ([]dto.Project, error) {
	listOfProjects, err := p.projectRepo.FindByFK(uuid2)
	if err != nil {
		return nil, err
	}

	projectsDto := gofp.Map(gofp.NewStreamFromSlice(listOfProjects), func(t model.Projects) dto.Project {
		pro := dto.Project{
			ProjectName: t.ProjectName,
			GithubLink:  t.GithubLink,
			IsActive:    t.IsActive,
		}
		return pro
	}).ToSlice()
	return projectsDto, nil
}

func NewProjectService(projectRepo repo.ProjectRepo) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (p projectService) NewProject(projects []dto.Project) error {
	projectModel := gofp.Map(gofp.NewStreamFromSlice(projects), func(t dto.Project) model.Projects {
		return model.Projects{ProjectName: t.ProjectName, GithubLink: t.GithubLink, IsActive: t.IsActive}
	}).ToSlice()

	return p.projectRepo.SaveAll(projectModel)

}

func (p projectService) UpdateProject(projects []dto.Project) error {
	projectModel := gofp.Map(gofp.NewStreamFromSlice(projects), func(t dto.Project) model.Projects {
		return model.Projects{ProjectName: t.ProjectName, GithubLink: t.GithubLink, IsActive: t.IsActive}
	}).ToSlice()

	return p.projectRepo.SaveAll(projectModel)
}

func (p projectService) DeleteProject(uuid2 uuid.UUID) error {
	return p.projectRepo.Delete(uuid2)
}
