package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/controllers"
)

type RouterInitializer struct {
	userController *controllers.UserController
}

func NewRouterInitializer(uc *controllers.UserController) *RouterInitializer {
	return &RouterInitializer{userController: uc}
}

func (r RouterInitializer) InitRouter() *gin.Engine {
	router := gin.Default()

	routerGroup := router.Group("/api")

	routerGroup.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "OK") })
	routerGroup.POST("/user", r.userController.CreateUser)

	return router
}
