package web

import "web-service-gin/model/dto"

type UserResponse struct {
	Id            int                       `json:"id"`
	FirstName     string                    `json:"firstName"`
	LastName      string                    `json:"lastName"`
	Email         string                    `json:"email"`
	Age           int                       `json:"age"`
	Subscriptions []dto.UserSubscriptionDto `json:"subscriptions"`
}
