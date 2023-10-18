package router

import (
	"recycle/app/middlewares"
	userController "recycle/features/user/controller"
	userRepository "recycle/features/user/repository"
	userUsecase "recycle/features/user/usecase"

	rubbishController "recycle/features/rubbish/controller"
	rubbishRepository "recycle/features/rubbish/repository"
	rubbishUsecase "recycle/features/rubbish/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// User
	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	userController := userController.NewUserControllers(userUsecase)

	// Rubbish
	rubbishRepository := rubbishRepository.NewRubbishRepository(db)
	rubbishUsecase := rubbishUsecase.NewRubbishUsecase(rubbishRepository)
	rubbishController := rubbishController.NewRubbishControllers(rubbishUsecase)

	// User & Admin CRUD
	user := e.Group("/users") 
	user.POST("/register", userController.CreateUser)
	user.POST("/login", userController.Login)
	user.GET("", userController.GetAllUser, middlewares.JWTMiddleware())
	user.GET("/:id", userController.GetUser, middlewares.JWTMiddleware())
	user.PUT("/:id", userController.Update, middlewares.JWTMiddleware())
	user.DELETE("/:id", userController.Delete, middlewares.JWTMiddleware())

	rubbish := e.Group("/rubbish")
	rubbish.POST("", rubbishController.CreateRubbish, middlewares.JWTMiddleware())
	rubbish.GET("", rubbishController.GetAllRubbish, middlewares.JWTMiddleware())
	rubbish.GET("/:id", rubbishController.GetRubbish, middlewares.JWTMiddleware())
	rubbish.PUT("/:id", rubbishController.UpdateRubbish, middlewares.JWTMiddleware())
	rubbish.DELETE("/:id", rubbishController.DeleteRubbish, middlewares.JWTMiddleware())
}
