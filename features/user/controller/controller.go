package controller

import (
	"net/http"
	"recycle/app/middlewares"
	"recycle/features/user"
	"recycle/helper"
	"strings"

	"github.com/google/uuid"
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
	dataInput := UserRequest{}

	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
	}

	data := RequestMain(dataInput)

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

func (uco *UserController) LoginUser(c echo.Context) error {
	var login UserLogin
	errBind := c.Bind(&login)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("eror bind"))
	}

	user, token, err := uco.userUseCase.CheckLogin(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("login failed"))
	}

	response := UserLoginResponse{
		Id:    user.Id,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login successful", response))
}

func (uco *UserController) GetUser(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
    if idToken == uuid.Nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }

    result, err := uco.userUseCase.GetById(idToken.String())

    if err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
    }

    var usersResponse = MainResponse(result)

    return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get profile", usersResponse))
}

func (uco *UserController) Update(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	userReq := UserRequest{}
	errBind := c.Bind(&userReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
	}

	userMain := RequestMain(userReq)
	data, err := uco.userUseCase.UpdateById(idToken.String(), userMain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", MainResponse(data)))
}

func (uco *UserController) Delete(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
	
	err := uco.userUseCase.DeleteById(idToken.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccesResponses("success delete user"))
}
