package routes

import (
	"service-pace11/controller"
	"service-pace11/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/notes", controller.NewNoteController(repository.NewNoteRepository()).GetNotes)
	api.GET("/note/:id", controller.NewNoteController(repository.NewNoteRepository()).GetNote)
	api.POST("/note", controller.NewNoteController(repository.NewNoteRepository()).CreateNote)
	api.PATCH("/note/:id", controller.NewNoteController(repository.NewNoteRepository()).UpdateNote)
	api.DELETE("/note/:id", controller.NewNoteController(repository.NewNoteRepository()).DeleteNote)
}
