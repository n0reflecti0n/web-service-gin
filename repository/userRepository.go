package repository

import (
	"gorm.io/gorm"
	"web-service-gin/model/dao"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) FindAll(users []dao.User) ([]dao.User, error) {
	result := ur.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (ur UserRepository) Save(user *dao.User) int {
	ur.db.Create(&user)
	return user.Id
}
