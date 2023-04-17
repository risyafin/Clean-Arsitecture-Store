package logins

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) Login(user User) error {
	result := repo.DB.Where("username = ?", user.Username).Where("password = ?", user.Password).First(&user)
	return result.Error
}
