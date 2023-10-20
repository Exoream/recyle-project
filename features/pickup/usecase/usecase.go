package usecase

import (
	"errors"
	"recycle/features/pickup/entity"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type pickupUseCase struct {
	pickupRepo entity.PickupDataInterface
	validate   *validator.Validate
}

func NewPickupUsecase(pickupRepo entity.PickupDataInterface) entity.UseCaseInterface {
	return &pickupUseCase{
		pickupRepo: pickupRepo,
		validate:   validator.New(),
	}
}

// Create implements entity.UseCaseInterface.
func (uc *pickupUseCase) Create(data entity.Main) error {
	errValidate := uc.validate.Struct(data)
	if errValidate != nil {
		return errValidate
	}

	data.Status = "schedule"

	datePattern := `^\d{4}-\d{2}-\d{2}$`
	matched, _ := regexp.MatchString(datePattern, data.PickupDate)
	if !matched {
		return errors.New("PickupDate should be in the format YYYY-MM-DD")
	}
	
	err := uc.pickupRepo.Create(data)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements entity.UseCaseInterface.
func (uc *pickupUseCase) DeleteById(id string) error {
	if id == "" {
		return errors.New("id not found")
	}

	data, err := uc.pickupRepo.GetById(id)
	if err != nil {
		return err
	}

	if data.Status == "done" {
		return errors.New("pickup is already completed and cannot be deleted")
	}

	err = uc.pickupRepo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateById implements entity.UseCaseInterface.
func (uc *pickupUseCase) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	if id == "" {
		return entity.Main{}, errors.New("id not found")
	}

	data, err = uc.pickupRepo.UpdateById(id, updated)
	if err != nil {
		return entity.Main{}, err
	}

	if data.Status == "done" {
		return entity.Main{}, errors.New("pickup is already completed and cannot be updated or deleted")
	}

	return data, nil
}

// GetById implements entity.UseCaseInterface.
func (uc *pickupUseCase) GetById(id string) (entity.Main, error) {
	if id == "" {
		return entity.Main{}, errors.New("invalid id")
	}

	idPickup, err := uc.pickupRepo.GetById(id)
	return idPickup, err
}

// FindAllPickup implements entity.UseCaseInterface.
func (uc *pickupUseCase) FindAllPickup() ([]entity.Main, error) {
	pickup, err := uc.pickupRepo.FindAllPickup()
	if err != nil {
		return nil, err
	}
	return pickup, nil
}
