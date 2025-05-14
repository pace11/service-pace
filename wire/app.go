package wire

import (
	"service-pace11/controllers"
	"service-pace11/repository"
)

type App struct {
	NoteController         *controllers.NoteController
	AuthController         *controllers.AuthController
	CommentController      *controllers.CommentController
	LikeController         *controllers.LikeController
	NotificationController *controllers.NotificationController
	RecipeController       *controllers.RecipeController
	UserController         *controllers.UserController
}

func InitApp() *App {
	// repository
	noteRepo := repository.NewNoteRepository()
	notificationRepo := repository.NewNotificationRepository()
	authRepo := repository.NewAuthRepository()
	commentRepo := repository.NewCommentRepository(notificationRepo)
	likeRepo := repository.NewLikeRepository(notificationRepo)
	recipeRepo := repository.NewRecipeRepository(notificationRepo)
	userRepo := repository.NewUserRepository()

	// controller
	noteCtrl := controllers.NewNoteController(noteRepo)
	notificationCtrl := controllers.NewNotificationController(notificationRepo)
	authCtrl := controllers.NewAuthController(authRepo)
	commentCtrl := controllers.NewCommentController(commentRepo)
	likeCtrl := controllers.NewLikeController(likeRepo)
	recipeCtrl := controllers.NewRecipeController(recipeRepo)
	userCtrl := controllers.NewUserController(userRepo)

	return &App{
		NoteController:         noteCtrl,
		NotificationController: notificationCtrl,
		AuthController:         authCtrl,
		CommentController:      commentCtrl,
		LikeController:         likeCtrl,
		RecipeController:       recipeCtrl,
		UserController:         userCtrl,
	}
}
