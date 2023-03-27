package service

import (
	"github.com/google/uuid"
	"projectsync/dto"
	"projectsync/model"
	"projectsync/repo"
	"projectsync/utils"
)

type UserService interface {
	Register(register dto.Register) error
	SignIn(in dto.SignIn) error
	UpdateUser(user dto.UpdateUser) error
	DeleteUser(uuid uuid.UUID) error
}

type userService struct {
	userRepo repo.UserdetailsRepo
}

func NewUserService(userRepo repo.UserdetailsRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(register dto.Register) error {
	if err := utils.Validate(register); err != nil {
		return err
	}
	return u.userRepo.Save(&model.UserDetails{})
}

func (u *userService) SignIn(in dto.SignIn) error {
	``
}

func (u *userService) UpdateUser(user dto.UpdateUser) error {
}

func (u *userService) DeleteUser(uuid uuid.UUID) error {
}
