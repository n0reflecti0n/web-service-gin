package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/dig"
	"log"
	"web-service-gin/config"
	"web-service-gin/controllers"
	"web-service-gin/mapper"
	"web-service-gin/repository"
	"web-service-gin/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	config.ConnectToDB()
}

func main() {

	container := dig.New()

	container.Provide(config.ConnectToDB)
	container.Provide(config.NewRouterInitializer)

	container.Provide(controllers.NewUserController)
	container.Provide(service.NewUserService)
	container.Provide(repository.NewUserRepository)
	container.Provide(mapper.NewUserMapper)

	// Invoke the handler and setup routes
	err := container.Invoke(func(r *config.RouterInitializer) {
		router := r.InitRouter()
		if err := router.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	})

	if err != nil {
		log.Fatal(err)
	}

}
