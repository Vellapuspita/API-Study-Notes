package main

import (
	"study_notes/config"
	"study_notes/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	routes.UsersRoutes(router)
	routes.StudyNotesRoutes(router)
	routes.TopicsRoutes(router)
	routes.CollabRoutes(router)

	router.Run(":3000")
}
