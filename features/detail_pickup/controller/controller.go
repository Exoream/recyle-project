package controller

import (
	"net/http"
	"recycle/app/middlewares"
	"recycle/features/detail_pickup/entity"
	"recycle/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type DetailPickupController struct {
	detailPickupUseCase entity.UseCaseInterface
}

func NewPickupControllers(uc entity.UseCaseInterface) *DetailPickupController {
	return &DetailPickupController{
		detailPickupUseCase: uc,
	}
}

func (uco *DetailPickupController) CreateDetailPickup(c echo.Context) error {
	// Mendapatkan role pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	// Periksa apakah pengguna adalah admin.
	if role == "admin" {
		// Bind data
		dataInput := DetailPickupRequest{}
		errBind := c.Bind(&dataInput)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
		}

		data := RequestDetailPickup(dataInput)

		errCreate := uco.detailPickupUseCase.Create(data)
		if errCreate != nil {
			if strings.Contains(errCreate.Error(), "validation") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
			} else {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data: " + errCreate.Error()))
			}
		}

		return c.JSON(http.StatusCreated, helper.SuccesResponses("success create data"))
	} else {
		// Jika pengguna bukan admin, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *DetailPickupController) GetAllDetailPickup(c echo.Context) error {
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

	if role == "admin" {
	// memanggil function dari usecase
	responseData, err := uco.detailPickupUseCase.FindAllDetailPickup()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
	}

	DetailPickupData := MapModelsToController(responseData)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all detail pickup", DetailPickupData))
	} else {
		// Jika bukan admin, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}