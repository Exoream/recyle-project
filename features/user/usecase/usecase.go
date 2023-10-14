package usecase

import (
	"recycle/features/user"

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
