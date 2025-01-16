package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/controller"
)

type RouterInitializer struct {
	userController *controller.UserController
}

func NewRouterInitializer(uc *controller.UserController) *RouterInitializer {
	return &RouterInitializer{userController: uc}
}

func (r RouterInitializer) InitRouter() *gin.Engine {
	router := gin.Default()

	routerGroup := router.Group("/api")

	routerGroup.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	routerGroup.GET("/user/:id", r.userController.FindUserById)
	routerGroup.POST("/user", r.userController.CreateUser)

	return router
}
