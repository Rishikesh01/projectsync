package model

import (
	"github.com/google/uuid"
	"time"
)

type Projects struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ProjectName string
	GithubLink  string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}
