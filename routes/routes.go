package routes

import (
	"service-pace11/controllers"
	"service-pace11/middlewares"
	"service-pace11/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	// notes
	api.GET("/notes", controllers.NewNoteController(repository.NewNoteRepository()).GetNotes)
	api.GET("/note/:id", controllers.NewNoteController(repository.NewNoteRepository()).GetNote)
	api.POST("/note", controllers.NewNoteController(repository.NewNoteRepository()).CreateNote)
	api.PATCH("/note/:id", controllers.NewNoteController(repository.NewNoteRepository()).UpdateNote)
	api.DELETE("/note/:id", controllers.NewNoteController(repository.NewNoteRepository()).DeleteNote)

	// auth (login, register, forgot-password)
	auth := api.Group("/auth")
	auth.POST("/login", controllers.NewLoginController(repository.NewLoginRepository()).Login)

	api.Use(middlewares.JWTAuthMiddleware())
	// recipes
	api.GET("/recipes", controllers.NewRecipeController(repository.NewRecipeRepository()).GetRecipes)
	api.GET("/recipe/:id", controllers.NewRecipeController(repository.NewRecipeRepository()).GetRecipe)
}
