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
	return u.userRepo.Save(&model.UserDetails{Email: register.Email, Password: register.Password, Name: register.Name})
}

func (u *userService) UpdateUser(user dto.UpdateUser) error {
	if err := utils.Validate(user); err != nil {
		return err
	}

	userDetails, err := u.userRepo.FindByEmail(user.Email)
	if err != nil {
		return err
	}

	userDetails.Email = user.Email
	userDetails.Password = user.Password
	return u.userRepo.Save(userDetails)

}

func (u *userService) DeleteUser(uid uuid.UUID) error {
	return u.userRepo.Delete(uid)
}
