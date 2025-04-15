package routes

import (
	"study_notes/controllers"

	"github.com/gin-gonic/gin"
)

func CollabRoutes(router *gin.Engine) {
	collab := router.Group("/collab")
	{
		collab.GET("", controllers.GetCollab)
		collab.POST("", controllers.CreateCollab)
		collab.GET("/:id", controllers.GetCollabByID)
		collab.PUT("/:id", controllers.UpdateCollab)
		collab.DELETE("/:id", controllers.DeleteCollab)
	}
}