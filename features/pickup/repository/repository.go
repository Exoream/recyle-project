package repository

import (
	"errors"
	"recycle/features/pickup/entity"
	"recycle/features/pickup/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type pickupRepository struct {
	db *gorm.DB
}

func NewPickupRepository(db *gorm.DB) entity.PickupDataInterface {
	return &pickupRepository{
		db: db,
	}
}

// Create implements entity.PickupDataInterface.
func (u *pickupRepository) Create(data entity.Main) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	dataInput := model.MapMainToModel(data)
	dataInput.Id = newUUID.String()

	tx := u.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteById implements entity.PickupDataInterface.
func (u *pickupRepository) DeleteById(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var pickup model.Pickup
	result := u.db.Where("id = ?", uuidID).Delete(&pickup)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete data, record not found")
	}

	return nil
}

// UpdateById implements entity.PickupDataInterface.
func (u *pickupRepository) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return entity.Main{}, err
	}

	var pickupData model.Pickup
	resultFind := u.db.First(&pickupData, uuidID)
	if resultFind.Error != nil {
		return entity.Main{}, resultFind.Error
	}

	u.db.Model(&pickupData).Updates(model.MapMainToModel(updated))

	data = model.MapModelToMain(pickupData)
	return data, nil
}

// GetById implements entity.PickupDataInterface.
func (u *pickupRepository) GetById(id string) (entity.Main, error) {
	var userData model.Pickup
	result := u.db.Where("id = ?", id).First(&userData)
	if result.Error != nil {
		return entity.Main{}, result.Error
	}

	var userById = model.MapModelToMain(userData)
	return userById, nil
}

// FindAllPickup implements entity.PickupDataInterface.
func (u *pickupRepository) FindAllPickup() ([]entity.Main, error) {
	var pickup []model.Pickup

	err := u.db.Find(&pickup).Error
	if err != nil {
		return nil, err
	}

	var allPickup []entity.Main = model.ModelToMainMapping(pickup)

	return allPickup, nil
}

// GetByStatus implements entity.PickupDataInterface.
func (u *pickupRepository) GetByStatus(status string) ([]entity.Main, error) {
	var pickupData []model.Pickup
    result := u.db.Where("status = ?", status).Find(&pickupData)
    if result.Error != nil {
        return nil, result.Error
    }

    var pickup []entity.Main
	pickup = model.ModelToMainMapping(pickupData)
    return pickup, nil
}