package service

import (
	"context"
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

func (ss SubscriptionService) GetUserSubscriptions(userId int) []dto.UserSubscriptionDto {
	ctxWithTimeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	responseChan := make(chan []dto.UserSubscriptionDto, 1)

	go func() {
		response, err := ss.fetchUserSubscriptions(userId)
		if err != nil {
			log.Println("Api call error:", err)
			responseChan <- []dto.UserSubscriptionDto{}
			return
		}

		responseChan <- response
	}()

	select {
	case response := <-responseChan:
		return response
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
