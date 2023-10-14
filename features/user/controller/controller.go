package controller

import (
	"net/http"
	"recycle/features/user"
	"recycle/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase user.UseCaseInterface
}

func NewUserControllers(uc user.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (uco *UserController) CreateUser(c echo.Context) error {
	dataInput := UserResponse{}

	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
	}

	data := user.Main{
		Id:          dataInput.ID,
		Name:        dataInput.Name,
		Email:       dataInput.Email,
		Password:    dataInput.Password,
		Gender:      dataInput.Gender,
		Age:         dataInput.Age,
		Address:     dataInput.Address,
		SaldoPoints: dataInput.SaldoPoints,
	}

	errCreate := uco.userUseCase.Create(data)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
		} else {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data"))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccesResponses("success create data"))
}
