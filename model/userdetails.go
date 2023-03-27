package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserDetails struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string
	Password  string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Projects  []Projects
}
