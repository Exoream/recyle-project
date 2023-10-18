package controller

import (
	"net/http"
	"recycle/app/middlewares"
	"recycle/features/user/entity"
	"recycle/helper"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase entity.UseCaseInterface
}

func NewUserControllers(uc entity.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (uco *UserController) CreateUser(c echo.Context) error {
	// Bind data
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

func (uco *UserController) Login(c echo.Context) error {
	// Bind data
	var login UserLogin
	errBind := c.Bind(&login)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("eror bind"))
	}

	// Memanggil func di usecase
	user, token, err := uco.userUseCase.CheckLogin(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("login failed"))
	}

    response := LoginResponse(user.Id, user.Email, token)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login successful", response))
}


func (uco *UserController) GetUser(c echo.Context) error {
    // Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
    
    // Mendapatkan email pengguna dari token
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

    idParamStr := c.Param("id")
    idParam, errId := uuid.Parse(idParamStr)
    if errId != nil {
        return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
    }

    // Periksa apakah ID dari token sama dengan ID dari parameter URL.
    if idToken.String() == idParam.String() || role == "admin" {
        // Jika sesuai, izinkan akses ke profil pengguna.
        result, err := uco.userUseCase.GetById(idParam.String())
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
        }

        var usersResponse = MainResponse(result)

        return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get profile", usersResponse))
    }

    // Jika tidak ada izin yang sesuai, kembalikan pesan "unauthorized."
    return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
}


func (uco *UserController) Update(c echo.Context) error {
    // Extra token from id
    idToken := middlewares.ExtractToken(c)
    if idToken == uuid.Nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }

    // Extract the user's role from the token
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

    idParamStr := c.Param("id")
    idParam, errId := uuid.Parse(idParamStr)
    if errId != nil {
        return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
    }

    
    userData, err := uco.userUseCase.GetById(idParam.String())
    if err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error retrieving user data"))
    }

    if idToken.String() == idParam.String() || role == "admin" {
        userReq := UserRequest{}
        errBind := c.Bind(&userReq)
        if errBind != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error binding data"))
        }

        userMain := RequestMain(userReq)

        if role != "admin" && userMain.SaldoPoints != userData.SaldoPoints {
            return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("You do not have permission to update 'saldo_points'"))
        }

        if role != "admin" && userMain.Role != userData.Role {
            return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("You do not have permission to update 'role'"))
        }

        data, err := uco.userUseCase.UpdateById(idParam.String(), userMain)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }

        return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", MainResponse(data)))
    } else {
        // If it's not an admin or the data owner, return an "unauthorized" response
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}


func (uco *UserController) Delete(c echo.Context) error {
    // Extra token dari id
    idToken := middlewares.ExtractToken(c)
    if idToken == uuid.Nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }

    // Mendapatkan email pengguna dari token
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

    idParamStr := c.Param("id")
    idParam, errId := uuid.Parse(idParamStr)
    if errId != nil {
        return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
    }

    if idToken.String() == idParam.String() || role == "admin" {
        err := uco.userUseCase.DeleteById(idParam.String())
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }

        return c.JSON(http.StatusOK, helper.SuccesResponses("success delete user"))
    } else {
        // Jika bukan admin dan bukan pemilik data, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}


func (uco *UserController) GetAllUser(c echo.Context) error {
    // Mendapatkan role pengguna dari token
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

    if role != "admin" {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }

	// memanggil function dari usecase
	responseData, err := uco.userUseCase.FindAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
	}

	userGetAllData := make([]UserResponse, 0)
	for _, value := range responseData {
		userResponse := MainResponse(value)
		userGetAllData = append(userGetAllData, userResponse)
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get all user data", userGetAllData))
}





