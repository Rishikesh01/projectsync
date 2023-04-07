package service

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/Rishikesh01/gofp"
	"github.com/google/uuid"
	"hash"
	"log"
	"projectsync/dto"
	"projectsync/model"
	"projectsync/repo"
)

type ProjectService interface {
	NewProject([]dto.NewProject) error
	UpdateProject([]dto.UpdateProject) error
	DeleteProject(uuid uuid.UUID) error
	ListProjects(uuid uuid.UUID) ([]dto.NewProject, error)
}

type projectService struct {
	projectRepo repo.ProjectRepo
	hash        hash.Hash
}

func NewProjectService(projectRepo repo.ProjectRepo) ProjectService {
	return &projectService{projectRepo: projectRepo, hash: sha256.New()}
}

func (p *projectService) NewProject(projects []dto.NewProject) error {
	projectModel := gofp.Map(gofp.NewStreamFromSlice(projects), func(t dto.NewProject) model.Projects {
		bytearray, err := json.Marshal(t)
		if err != nil {
			log.Println(err)
		}
		hashString := hex.EncodeToString(p.hash.Sum(bytearray))
		return model.Projects{ProjectName: t.ProjectName, GithubLink: t.GithubLink, IsActive: t.IsActive, Hash: hashString}
	}).ToSlice()

	return p.projectRepo.SaveAll(projectModel)

}

func (p *projectService) ListProjects(uuid2 uuid.UUID) ([]dto.NewProject, error) {
	listOfProjects, err := p.projectRepo.FindByFK(uuid2)
	if err != nil {
		return nil, err
	}

	projectsDto := gofp.Map(gofp.NewStreamFromSlice(listOfProjects), func(t model.Projects) dto.NewProject {
		pro := dto.NewProject{
			ProjectName: t.ProjectName,
			GithubLink:  t.GithubLink,
			IsActive:    t.IsActive,
		}
		return pro
	}).ToSlice()
	return projectsDto, nil
}

func (p *projectService) UpdateProject(projects []dto.UpdateProject) error {
	var projectModel []model.Projects
	for _, val := range projects {
		m, err := p.projectRepo.FindByID(val.ID)
		if err != nil {
			return err
		}
		bytearray, err := json.Marshal(val.Project)
		if err != nil {
			log.Println(err)
		}
		hashString := hex.EncodeToString(p.hash.Sum(bytearray))
		if m.Hash != hashString {
			projectModel = append(projectModel, model.Projects{ProjectName: val.Project.ProjectName, GithubLink: val.Project.GithubLink, IsActive: val.Project.IsActive, Hash: hashString})
		}
	}

	return p.projectRepo.SaveAll(projectModel)
}

func (p *projectService) DeleteProject(uuid2 uuid.UUID) error {
	return p.projectRepo.Delete(uuid2)
}
