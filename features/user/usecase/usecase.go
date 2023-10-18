package usecase

import (
	"errors"
	"recycle/app/middlewares"
	"recycle/features/user/entity"
	"recycle/helper"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type userUseCase struct {
	userRepo entity.UserDataInterface
	validate *validator.Validate
}

func NewUserUsecase(userRepo entity.UserDataInterface) entity.UseCaseInterface {
	return &userUseCase{
		userRepo: userRepo,
		validate: validator.New(),
	}
}

// Create implements user.UseCaseInterface.
func (uc *userUseCase) Create(data entity.Main) error {
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
func (uc *userUseCase) CheckLogin(email string, password string) (entity.Main, string, error) {
	if email == "" || password == "" {
		return entity.Main{}, "", errors.New("error validation: email dan password harus diisi")
	}

	dataLogin, err := uc.userRepo.CheckLogin(email, password)

	if err != nil {
		return entity.Main{}, "", err
	}

	id, err := uuid.Parse(dataLogin.Id)
	if err != nil {
		return entity.Main{}, "", err
	}

	if helper.CheckPasswordHash(dataLogin.Password, password) {
		token, err := middlewares.CreateToken(id, dataLogin.Role)
		if err != nil {
			return entity.Main{}, "", err
		}

		return dataLogin, token, nil
	}

	return entity.Main{}, "", errors.New("Login failed")
}

// GetById implements user.UseCaseInterface.
func (uc *userUseCase) GetById(id string) (entity.Main, error) {
	if id == "" {
		return entity.Main{}, errors.New("invalid id")
	}

	idUser, err := uc.userRepo.GetById(id)
	return idUser, err
}

// UpdateById implements user.UseCaseInterface.
func (uc *userUseCase) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	if id == "" {
		return entity.Main{}, errors.New("id not found")
	}

	if updated.Password != "" {
		hash, hashErr := helper.HashPassword(updated.Password)
		if hashErr != nil {
			return entity.Main{}, hashErr
		}
		updated.Password = hash
	}

	data, err = uc.userRepo.UpdateById(id, updated)
	if err != nil {
		return entity.Main{}, err
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

// FindAllUsers implements user.UseCaseInterface.
func (uc *userUseCase) FindAllUsers() ([]entity.Main, error) {
	users, err := uc.userRepo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}