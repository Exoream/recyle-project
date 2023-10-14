package router

import (
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

	e.POST("/users/register", userController.CreateUser)
}
