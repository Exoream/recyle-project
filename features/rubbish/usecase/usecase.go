package usecase

import (
	"errors"
	"recycle/features/rubbish/entity"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	rubbishRepo entity.RubbishDataInterface
	validate    *validator.Validate
}

func NewRubbishUsecase(rubbishRepo entity.RubbishDataInterface) entity.UseCaseInterface {
	return &userUseCase{
		rubbishRepo: rubbishRepo,
		validate:    validator.New(),
	}
}

// Create implements rubbish.UseCaseInterface.
func (uc *userUseCase) Create(data entity.Main) error {
	errValidate := uc.validate.Struct(data)
	if errValidate != nil {
		return errValidate
	}

	err := uc.rubbishRepo.Create(data)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements rubbish.UseCaseInterface.
func (uc *userUseCase) DeleteById(id string) error {
	if id == "" {
		return errors.New("id not found")
	}

	err := uc.rubbishRepo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

// FindAllUsers implements rubbish.UseCaseInterface.
func (uc *userUseCase) FindAllRubbish() ([]entity.Main, error) {
	rubbish, err := uc.rubbishRepo.FindAllRubbish()
	if err != nil {
		return nil, err
	}
	return rubbish, nil
}

// GetById implements rubbish.UseCaseInterface.
func (uc *userUseCase) GetById(id string) (entity.Main, error) {
	if id == "" {
		return entity.Main{}, errors.New("invalid id")
	}

	idRubbish, err := uc.rubbishRepo.GetById(id)
	return idRubbish, err
}

// UpdateById implements rubbish.UseCaseInterface.
func (uc *userUseCase) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	if id == "" {
		return entity.Main{}, errors.New("id not found")
	}

	data, err = uc.rubbishRepo.UpdateById(id, updated)
	if err != nil {
		return entity.Main{}, err
	}
	return data, nil
}

// GetByType implements entity.UseCaseInterface.
func (uc *userUseCase) GetByType(typeRubbish string) ([]entity.Main, error) {
	if typeRubbish == "" {
		return nil, errors.New("typeRubbish parameter is required")
	}

	typeName, err := uc.rubbishRepo.GetByType(typeRubbish)
	return typeName, err
}