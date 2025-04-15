package routes

import (
	"study_notes/controllers"

	"github.com/gin-gonic/gin"
)

func TopicsRoutes(router *gin.Engine) {
	topics := router.Group("/topics")
	{
		topics.GET("", controllers.GetTopics)
		topics.POST("", controllers.CreateTopics)
		topics.GET("/:id", controllers.GetTopicsByID)
		topics.PUT("/:id", controllers.UpdateTopics)
		topics.DELETE("/:id", controllers.DeleteTopics)
	}
}