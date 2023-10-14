package usecase

import (
	"errors"
	"recycle/features/user"
	// "recycle/helper"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	userRepo user.UserDataInterface
	validate *validator.Validate
}

func NewUserUsecase(userRepo user.UserDataInterface) user.UseCaseInterface {
	return &userUseCase{
		userRepo: userRepo,
		validate: validator.New(),
	}
}

func (uc *userUseCase) Create(data user.Main) error {
	errValidate := uc.validate.Struct(data)
	if errValidate != nil {
		return errValidate
	}

	err := uc.userRepo.Create(data)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) CheckLogin(email string, password string) (user.Main, string, error) {
	if email == "" || password == "" {
		return user.Main{}, "", errors.New("error validation: email dan password harus diisi")
	}

	// data, err := uc.userRepo.GetByEmail(email)
    // if err != nil {
    //     return user.Main{}, "", err
    // }


	dataLogin, token, err := uc.userRepo.CheckLogin(email, password)
	return dataLogin, token, err
}
