package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
	"web-service-gin/model/dto"
)

type SubscriptionService struct {
}

func NewSubscriptionService() *SubscriptionService {
	return &SubscriptionService{}
}

func (ss SubscriptionService) GetUserSubscriptions(c *gin.Context, userId int) []dto.UserSubscriptionDto {
	ctxWithTimeout, cancelFunc := context.WithTimeout(c, 5*time.Second)
	defer cancelFunc()

	responseChan := make(chan struct {
		data []dto.UserSubscriptionDto
		err  error
	}, 1)

	go func() {
		response, err := ss.fetchUserSubscriptions(userId)
		if err != nil {
			log.Println("Api call error:", err)
		}

		responseChan <- struct {
			data []dto.UserSubscriptionDto
			err  error
		}{response, err}
	}()

	select {

	case response := <-responseChan:
		if response.err != nil {
			return []dto.UserSubscriptionDto{}
		}
		return response.data

	case <-ctxWithTimeout.Done():
		log.Println("Api call timeout")
		return []dto.UserSubscriptionDto{}

	}
}

func (ss SubscriptionService) fetchUserSubscriptions(userId int) ([]dto.UserSubscriptionDto, error) {
	log.Printf("Initiating user subscriptions request for userId=%d", userId)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	return []dto.UserSubscriptionDto{
		{
			Id:   1,
			Name: "SampleSubscription",
		},
	}, nil
}
