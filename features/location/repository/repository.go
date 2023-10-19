package repository

import (
	"errors"
	"recycle/features/location/entity"
	"recycle/features/location/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) entity.LocationDataInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements entity.LocationDataInterface.
func (u *userRepository) Create(data entity.Main) error {
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

// DeleteById implements entity.LocationDataInterface.
func (u *userRepository) DeleteById(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var location model.Location
	result := u.db.Where("id = ?", uuidID).Delete(&location)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete data, record not found")
	}

	return nil
}

// FindAllLocation implements entity.LocationDataInterface.
func (u *userRepository) FindAllLocation() ([]entity.Main, error) {
	var location []model.Location

	err := u.db.Find(&location).Error
	if err != nil {
		return nil, err
	}

	var allUser []entity.Main = model.ModelToMainMapping(location)

	return allUser, nil
}

// GetById implements entity.LocationDataInterface.
func (u *userRepository) GetById(id string) (entity.Main, error) {
	var locationData model.Location
	result := u.db.Where("id = ?", id).First(&locationData)
	if result.Error != nil {
		return entity.Main{}, result.Error
	}

	var userById = model.MapModelToMain(locationData)
	return userById, nil
}

// GetByCity implements entity.LocationDataInterface.
func (u *userRepository) GetByCity(city string) ([]entity.Main, error) {
	var locationData []model.Location
    result := u.db.Where("city = ?", city).Find(&locationData)
    if result.Error != nil {
        return nil, result.Error
    }

    var location []entity.Main
	location = model.ModelToMainMapping(locationData)
    return location, nil
}

// UpdateById implements entity.LocationDataInterface.
func (u *userRepository) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return entity.Main{}, err
	}

	var locationData model.Location
	resultFind := u.db.First(&locationData, uuidID)
	if resultFind.Error != nil {
		return entity.Main{}, resultFind.Error
	}

	u.db.Model(&locationData).Updates(model.MapMainToModel(updated))

	data = model.MapModelToMain(locationData)
	return data, nil
}
