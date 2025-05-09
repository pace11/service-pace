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

	// auth
	auth := api.Group("/auth")
	auth.POST("/login", controllers.NewAuthController(repository.NewAuthRepository()).Login)
	auth.POST("/register", controllers.NewAuthController(repository.NewAuthRepository()).Register)

	api.Use(middlewares.JWTAuthMiddleware())

	// user
	user := api.Group("/user")
	user.GET("/me", controllers.NewUserController(repository.NewUserRepository()).GetMe)
	user.POST("/update", controllers.NewUserController(repository.NewUserRepository()).Update)

	// recipes
	api.GET("/recipes", controllers.NewRecipeController(repository.NewRecipeRepository(repository.NewNotificationRepository())).GetRecipes)
	api.GET("/recipe/:id", controllers.NewRecipeController(repository.NewRecipeRepository(repository.NewNotificationRepository())).GetRecipe)
	api.POST("/recipe", controllers.NewRecipeController(repository.NewRecipeRepository(repository.NewNotificationRepository())).CreateRecipe)
	api.PATCH("/recipe/:id", controllers.NewRecipeController(repository.NewRecipeRepository(repository.NewNotificationRepository())).UpdateRecipe)
	api.DELETE("/recipe/:id", controllers.NewRecipeController(repository.NewRecipeRepository(repository.NewNotificationRepository())).DeleteRecipe)
	api.GET("/recipe/saves", controllers.NewRecipeController(repository.NewRecipeRepository(repository.NewNotificationRepository())).SavedRecipeIndex)
	api.POST("/recipe/save/:id", controllers.NewRecipeController(repository.NewRecipeRepository(repository.NewNotificationRepository())).SavedRecipe)

	// likes
	api.GET("/likes/recipe/:id", controllers.NewLikeController(repository.NewLikeRepository(repository.NewNotificationRepository())).GetLikesByRecipe)
	api.POST("/like/recipe/:id", controllers.NewLikeController(repository.NewLikeRepository(repository.NewNotificationRepository())).LikeByRecipe)
	api.POST("/unlike/recipe/:id", controllers.NewLikeController(repository.NewLikeRepository(repository.NewNotificationRepository())).UnlikeByRecipe)

	// comments
	api.GET("/comments/recipe/:id", controllers.NewCommentController(repository.NewCommentRepository(repository.NewNotificationRepository())).GetCommentsByRecipe)
	api.POST("/comment/recipe/:id", controllers.NewCommentController(repository.NewCommentRepository(repository.NewNotificationRepository())).CommentByRecipe)
	api.DELETE("/comment/:id", controllers.NewCommentController(repository.NewCommentRepository(repository.NewNotificationRepository())).DeleteComment)

	// notifications
	api.GET("/notifications", controllers.NewNotificationController(repository.NewNotificationRepository()).GetNotifications)
}
