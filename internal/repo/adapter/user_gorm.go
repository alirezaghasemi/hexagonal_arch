package adapter

import (
	"github.com/alirezaghasemi/hexagonal_arch/internal/entity"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) *GormUserRepo {
	return &GormUserRepo{db: db}
}

func (r *GormUserRepo) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *GormUserRepo) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
