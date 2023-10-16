package usecase

import (
	"errors"
	"recycle/app/middlewares"
	"recycle/features/user"
	"recycle/helper"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

// Create implements user.UseCaseInterface.
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

// CheckLogin implements user.UseCaseInterface.
func (uc *userUseCase) CheckLogin(email string, password string) (user.Main, string, error) {
	if email == "" || password == "" {
		return user.Main{}, "", errors.New("error validation: email dan password harus diisi")
	}

	dataLogin, err := uc.userRepo.CheckLogin(email, password)

	if err != nil {
		return user.Main{}, "", err
	}

	id, err := uuid.Parse(dataLogin.Id)
	if err != nil {
		return user.Main{}, "", err
	}

	if helper.CheckPasswordHash(dataLogin.Password, password) {
		token, err := middlewares.CreateToken(id)
		if err != nil {
			return user.Main{}, "", err
		}

		return dataLogin, token, nil
	}

	return user.Main{}, "", errors.New("Login failed")
}

// GetById implements user.UseCaseInterface.
func (uc *userUseCase) GetById(id string) (user.Main, error) {
	if id == "" {
		return user.Main{}, errors.New("invalid id")
	}

	idUser, err := uc.userRepo.GetById(id)
	return idUser, err
}

// UpdateById implements user.UseCaseInterface.
func (uc *userUseCase) UpdateById(id string, updated user.Main) (data user.Main, err error) {
	if id == "" {
		return user.Main{}, errors.New("id not found")
	}

	if updated.Password != "" {
		hash, hashErr := helper.HashPassword(updated.Password)
		if hashErr != nil {
			return user.Main{}, hashErr
		}
		updated.Password = hash
	}

	data, err = uc.userRepo.UpdateById(id, updated)
	if err != nil {
		return user.Main{}, err
	}
	return data, nil
}

// Delete implements user.UseCaseInterface.
func (uc *userUseCase) DeleteById(id string) error {
	if id == "" {
		return errors.New("id not found")
	}

	err := uc.userRepo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil

}
