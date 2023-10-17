package router

import (
	"recycle/app/middlewares"
	"recycle/features/user/controller"
	"recycle/features/user/repository"
	"recycle/features/user/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserControllers(userUsecase)

	// User & Admin CRUD 
	e.POST("/users/register", userController.CreateUser)
	e.POST("/users/login", userController.Login)
	e.GET("/users", userController.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/:id", userController.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users/:id", userController.Update, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.Delete, middlewares.JWTMiddleware())
}
