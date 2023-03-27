package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"projectsync/model"
)

type UserdetailsRepo interface {
	Save(model *model.UserDetails) error
	FindByEmail(email string) (*model.UserDetails, error)
	Delete(uuid uuid.UUID) error
}

type userDetailsRepo struct {
	db *gorm.DB
}

func NewUserdetailsRepo(db *gorm.DB) UserdetailsRepo {
	return &userDetailsRepo{db: db}
}
func (u userDetailsRepo) Save(model *model.UserDetails) error {
	return u.db.Save(model).Error
}

func (u userDetailsRepo) FindByEmail(email string) (*model.UserDetails, error) {
	var m model.UserDetails
	if err := u.db.Where("email=?", email).Find(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (u userDetailsRepo) Delete(uuid2 uuid.UUID) error {
	return u.db.Where("id=?", uuid2).Delete(&model.UserDetails{}).Error
}
