package service

import (
	"web-service-gin/mapper"
	"web-service-gin/model/web"
	"web-service-gin/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
	userMapper     *mapper.UserMapper
}

func NewUserService(ur *repository.UserRepository, um *mapper.UserMapper) *UserService {
	return &UserService{userRepository: ur, userMapper: um}
}

func (us UserService) CreateUser(request web.UserRequest) int {
	user := us.userMapper.ToUser(request)

	return us.userRepository.Save(&user)
}

func (us UserService) FindById(id int) (web.UserResponse, error) {
	var response web.UserResponse

	user, err := us.userRepository.FindById(int64(id))
	if err != nil {
		return response, err
	}

	response = us.userMapper.ToUserResponse(user)

	return response, nil
}
