package mapper

import (
	"time"
	"web-service-gin/model/dao"
	"web-service-gin/model/dto"
	"web-service-gin/model/web"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (um UserMapper) ToUser(request web.UserRequest) dao.User {
	return dao.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		Age:         request.Age,
		CreatedDate: time.Now(),
	}
}

func (um UserMapper) ToUserResponse(user dao.User, subscriptions []dto.UserSubscriptionDto) web.UserResponse {
	return web.UserResponse{
		Id:            user.Id,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Age:           user.Age,
		Subscriptions: subscriptions,
	}
}
