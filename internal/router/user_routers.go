package router

import (
	"task-manager/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, h handler.UserHandler) {

	userGroup := r.Group("/users")
	{
		userGroup.POST("/signup", h.Signup)
		// userGroup.POST("/login", handler.Login)
		// userGroup.GET("/:id", handler.GetUser)
		// userGroup.PUT("/:id", handler.UpdateUser)
		// userGroup.DELETE("/:id", handler.DeleteUser)

	}
}
