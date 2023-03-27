package service

import (
	"github.com/google/uuid"
	"projectsync/dto"
)

type UserService interface {
	Register(register dto.Register) error
	SignIn(in dto.SignIn) error
	UpdateUser(user dto.UpdateUser) error
	DeleteUser(uuid uuid.UUID) error
}
