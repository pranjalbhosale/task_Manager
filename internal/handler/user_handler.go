package handler

import (
	"net/http"
	"task-manager/internal/entity"
	"task-manager/internal/service"
	"task-manager/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Signup(c *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{service: s}
}

func (h *userHandler) Signup(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Log.Error("Invalid input: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid input",
			"body":    err.Error(),
		})
		return
	}

	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		logger.Log.Error("Failed to create user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create user",
			"body":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User created successfuly",
		"body":    createdUser,
	})

}
