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

	// User CRUD
	e.POST("/users/register", userController.CreateUser)
	e.POST("/users/login", userController.LoginUser)
	e.GET("/users", userController.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userController.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userController.Delete, middlewares.JWTMiddleware())
}
