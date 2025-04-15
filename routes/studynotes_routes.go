package routes

import (
	"study_notes/controllers"

	"github.com/gin-gonic/gin"
)

func StudyNotesRoutes(r *gin.Engine) {
	studyNoteRoutes := r.Group("/studynotes")
	{
		studyNoteRoutes.GET("", controllers.GetStudyNotes)
		studyNoteRoutes.POST("", controllers.CreateStudyNote)
		studyNoteRoutes.GET("/:id", controllers.GetStudyNoteByID)
		studyNoteRoutes.PUT("/:id", controllers.UpdateStudyNote)
		studyNoteRoutes.DELETE("/:id", controllers.DeleteStudyNote)
	}
}
