package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/dig"
	"log"
	"web-service-gin/config"
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

	config.PerformDependencyInjection(container)

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
