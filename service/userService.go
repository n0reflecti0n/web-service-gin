package service

import (
	"web-service-gin/mapper"
	"web-service-gin/model/web"
	"web-service-gin/repository"
)

type UserService struct {
	userRepository      *repository.UserRepository
	userMapper          *mapper.UserMapper
	subscriptionService *SubscriptionService
}

func NewUserService(ur *repository.UserRepository, um *mapper.UserMapper, ss *SubscriptionService) *UserService {
	return &UserService{userRepository: ur, userMapper: um, subscriptionService: ss}
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

	subscriptions := us.subscriptionService.GetUserSubscriptions(id)

	response = us.userMapper.ToUserResponse(user, subscriptions)

	return response, nil
}
