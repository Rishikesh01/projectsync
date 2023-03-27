package repo

import (
	"gorm.io/gorm"
	"projectsync/model"
)

type UserdetailsRepo interface {
	Save(model *model.UserDetails) error
	FindByEmail(email string) (*model.UserDetails, error)
}

type userDetailsRepo struct {
	db *gorm.DB
}

func NewUserdetailsRepo(db *gorm.DB) UserdetailsRepo {
	return &userDetailsRepo{db: db}
}
func (u userDetailsRepo) Save(model *model.UserDetails) error {
	//TODO implement me
	panic("implement me")
}

func (u userDetailsRepo) FindByEmail(email string) (*model.UserDetails, error) {
	//TODO implement me
	panic("implement me")
}
