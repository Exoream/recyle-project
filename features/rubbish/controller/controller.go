package controller

import (
	"net/http"
	"recycle/app/middlewares"
	"recycle/features/rubbish"
	"recycle/helper"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RubbishController struct {
	rubbishUseCase rubbish.UseCaseInterface
}

func NewRubbishControllers(uc rubbish.UseCaseInterface) *RubbishController {
	return &RubbishController{
		rubbishUseCase: uc,
	}
}

func (uco *RubbishController) CreateRubbish(c echo.Context) error {
    // Mendapatkan email pengguna dari token
    userEmail, err := middlewares.ExtractUserEmail(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting user email"))
    }

    // Periksa apakah pengguna adalah admin.
    if userEmail == "admin@gmail.com" {
        // Bind data
        dataInput := RubbishRequest{}
        errBind := c.Bind(&dataInput)
        if errBind != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
        }

        data := RequestMain(dataInput)

        errCreate := uco.rubbishUseCase.Create(data)
        if errCreate != nil {
            if strings.Contains(errCreate.Error(), "validation") {
                return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
            } else {
                return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data"))
            }
        }

        return c.JSON(http.StatusCreated, helper.SuccesResponses("success create data"))
    } else {
        // Jika pengguna bukan admin, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}


func (uco *RubbishController) GetRubbish(c echo.Context) error {
	// Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	idParamStr := c.Param("id")
	idParam, errId := uuid.Parse(idParamStr)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
	}

	// Izinkan akses ke data sampah (rubbish) berdasarkan ID tanpa pemeriksaan tambahan.
	result, err := uco.rubbishUseCase.GetById(idParam.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
	}

	var rubbishResponse = MainResponse(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get rubbish", rubbishResponse))
}

func (uco *RubbishController) UpdateRubbish(c echo.Context) error {
    // Mendapatkan email pengguna dari token
    userEmail, err := middlewares.ExtractUserEmail(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting user email"))
    }

    // Periksa apakah pengguna adalah admin.
    if userEmail == "admin@gmail.com" {
        idParamStr := c.Param("id")
        idParam, errId := uuid.Parse(idParamStr)
        if errId != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
        }

        rubbishReq := RubbishRequest{}
        errBind := c.Bind(&rubbishReq)
        if errBind != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
        }

        userMain := RequestMain(rubbishReq)
        data, err := uco.rubbishUseCase.UpdateById(idParam.String(), userMain)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }

        return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", MainResponse(data)))
    } else {
        // Jika bukan admin, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}

func (uco *RubbishController) DeleteRubbish(c echo.Context) error {
    // Mendapatkan email pengguna dari token
    userEmail, err := middlewares.ExtractUserEmail(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting user email"))
    }

    // Periksa apakah pengguna adalah admin.
    if userEmail == "admin@gmail.com" {
        idParamStr := c.Param("id")
        idParam, errId := uuid.Parse(idParamStr)
        if errId != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
        }

        err := uco.rubbishUseCase.DeleteById(idParam.String())
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }
        return c.JSON(http.StatusOK, helper.SuccesResponses("success delete rubbish"))
    } else {
        // Jika bukan admin, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}

func (uco *RubbishController) GetAllRubbish(c echo.Context) error {
    // Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// memanggil function dari usecase
	responseData, err := uco.rubbishUseCase.FindAllRubbish()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
	}

	rubbishGetAllData := make([]RubbishRespon, 0)
	for _, value := range responseData {
		userResponse := MainResponse(value)
		rubbishGetAllData = append(rubbishGetAllData, userResponse)
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get all user rubbish", rubbishGetAllData))
}



