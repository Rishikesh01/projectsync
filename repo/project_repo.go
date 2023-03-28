package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectsync/model"
)

type ProjectRepo interface {
	Save(project *model.Projects) error
	SaveAll(project []model.Projects) error
	FindByID(uuid uuid.UUID) (*model.Projects, error)
	FindByFK(uuid uuid.UUID) ([]model.Projects, error)
	Delete(uuid2 uuid.UUID) error
}

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) ProjectRepo {
	return &projectRepo{
		db: db,
	}
}

func (p *projectRepo) FindByFK(uuid2 uuid.UUID) ([]model.Projects, error) {
	var projects []model.Projects
	if err := p.db.Where("user_id=?", uuid2).Find(&projects).Error; err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return projects, nil
}

func (p *projectRepo) Delete(uuid2 uuid.UUID) error {
	return p.db.Where("id=?", uuid2).Delete(&model.Projects{}).Error
}

func (p *projectRepo) Save(project *model.Projects) error {
	return p.db.Save(project).Error
}

func (p *projectRepo) SaveAll(project []model.Projects) error {
	return p.db.Save(project).Error
}

func (p *projectRepo) FindByID(uid uuid.UUID) (*model.Projects, error) {
	var m model.Projects
	if err := p.db.Where("id=?", uid).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
