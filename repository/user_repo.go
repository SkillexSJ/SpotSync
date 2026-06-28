package repository

import (
	"github.com/SkillexSJ/SpotSync/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// create repo
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// create user
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByEmail
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
