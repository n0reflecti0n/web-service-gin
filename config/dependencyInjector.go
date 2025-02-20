package config

import (
	"go.uber.org/dig"
	"web-service-gin/controller"
	"web-service-gin/mapper"
	"web-service-gin/repository"
	"web-service-gin/service"
)

func PerformDependencyInjection(container *dig.Container) {
	container.Provide(ConnectToDB)
	container.Provide(NewRouterInitializer)

	container.Provide(controller.NewUserController)

	container.Provide(repository.NewUserRepository)

	container.Provide(mapper.NewUserMapper)

	container.Provide(service.NewUserService)
	container.Provide(service.NewSubscriptionService)
}
