package mapper

import (
	"time"
	"web-service-gin/model/dao"
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
