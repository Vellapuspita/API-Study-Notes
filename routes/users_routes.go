package routes

import (
	"study_notes/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("", controllers.GetUsers)
		users.POST("", controllers.CreateUser)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUsers)
		users.DELETE("/:id", controllers.DeleteUsers)
	}
}