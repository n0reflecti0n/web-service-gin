package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web-service-gin/model/web"
	"web-service-gin/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc UserController) CreateUser(c *gin.Context) {
	log.Println("User created")
	var user web.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := uc.userService.CreateUser(user)

	c.JSON(http.StatusOK, gin.H{"id": userId})
}
