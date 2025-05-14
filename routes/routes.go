package routes

import (
	"service-pace11/middlewares"
	"service-pace11/wire"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, app *wire.App) {
	api := r.Group("/api")
	// notes
	api.GET("/notes", app.NoteController.GetNotes)
	api.GET("/note/:id", app.NoteController.GetNote)
	api.POST("/note", app.NoteController.CreateNote)
	api.PATCH("/note/:id", app.NoteController.UpdateNote)
	api.DELETE("/note/:id", app.NoteController.DeleteNote)

	// auth
	auth := api.Group("/auth")
	auth.POST("/login", app.AuthController.Login)
	auth.POST("/register", app.AuthController.Register)

	api.Use(middlewares.JWTAuthMiddleware())

	// user
	user := api.Group("/user")
	user.GET("/me", app.UserController.GetMe)
	user.POST("/update", app.UserController.Update)

	// recipes
	api.GET("/recipes", app.RecipeController.GetRecipes)
	api.GET("/recipe/:id", app.RecipeController.GetRecipe)
	api.POST("/recipe", app.RecipeController.CreateRecipe)
	api.PATCH("/recipe/:id", app.RecipeController.UpdateRecipe)
	api.DELETE("/recipe/:id", app.RecipeController.DeleteRecipe)
	api.GET("/recipe/saves", app.RecipeController.SavedRecipeIndex)
	api.POST("/recipe/save/:id", app.RecipeController.SavedRecipe)

	// likes
	api.GET("/likes/recipe/:id", app.LikeController.GetLikesByRecipe)
	api.POST("/like/recipe/:id", app.LikeController.LikeByRecipe)
	api.POST("/unlike/recipe/:id", app.LikeController.UnlikeByRecipe)

	// comments
	api.GET("/comments/recipe/:id", app.CommentController.GetCommentsByRecipe)
	api.POST("/comment/recipe/:id", app.CommentController.CommentByRecipe)
	api.DELETE("/comment/:id", app.CommentController.DeleteComment)

	// notifications
	api.GET("/notifications", app.NotificationController.GetNotifications)
}
