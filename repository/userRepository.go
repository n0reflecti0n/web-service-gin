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

func (ur UserRepository) FindById(userId int64) (dao.User, error) {
	var user dao.User
	result := ur.db.First(&user, userId)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (ur UserRepository) Save(user *dao.User) int {
	ur.db.Create(&user)
	return user.Id
}
