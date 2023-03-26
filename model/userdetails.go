package model

import (
	"github.com/google/uuid"
	"time"
)

type UserDetails struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string
	Password  string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
