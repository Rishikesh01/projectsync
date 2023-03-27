package repo

import (
	"github.com/google/uuid"
	"projectsync/model"
)

type ProjectRepo interface {
	Save(project *model.Projects) error
	FindByID(uuid uuid.UUID) (*model.Projects, error)
}
