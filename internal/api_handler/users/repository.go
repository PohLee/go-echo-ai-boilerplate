package users

import (
	"github.com/PohLee/go-echo-ai-boilerplate/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}

type repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *model.User) error {
	if r.db == nil {
		return gorm.ErrInvalidDB
	}
	return r.db.Create(user).Error
}

func (r *repository) FindByEmail(email string) (*model.User, error) {
	if r.db == nil {
		return nil, gorm.ErrInvalidDB
	}
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByID(id uint) (*model.User, error) {
	if r.db == nil {
		return nil, gorm.ErrInvalidDB
	}
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
