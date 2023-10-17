package repository

import (
	"errors"
	"recycle/features/rubbish"
	"recycle/features/rubbish/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRubbishRepository(db *gorm.DB) rubbish.RubbishDataInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements rubbish.RubbishDataInterface.
func (u *userRepository) Create(data rubbish.Main) error {
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

// DeleteById implements rubbish.RubbishDataInterface.
func (u *userRepository) DeleteById(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var rubbish model.Rubbish
	result := u.db.Where("id = ?", uuidID).Delete(&rubbish)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete data, record not found")
	}

	return nil
}

// FindAllUsers implements rubbish.RubbishDataInterface.
func (u *userRepository) FindAllRubbish() ([]rubbish.Main, error) {
	var trash []model.Rubbish

	err := u.db.Find(&trash).Error
	if err != nil {
		return nil, err
	}

	var allUser []rubbish.Main = model.ModelToMainMapping(trash)

	return allUser, nil
}

// GetById implements rubbish.RubbishDataInterface.
func (u *userRepository) GetById(id string) (rubbish.Main, error) {
	var rubbishData model.Rubbish
	result := u.db.Where("id = ?", id).First(&rubbishData)
	if result.Error != nil {
		return rubbish.Main{}, result.Error
	}

	var userById = model.MapModelToMain(rubbishData)
	return userById, nil
}

// UpdateById implements rubbish.RubbishDataInterface.
func (u *userRepository) UpdateById(id string, updated rubbish.Main) (data rubbish.Main, err error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return rubbish.Main{}, err
	}

	var rubbishData model.Rubbish
	resultFind := u.db.First(&rubbishData, uuidID)
	if resultFind.Error != nil {
		return rubbish.Main{}, resultFind.Error
	}

	u.db.Model(&rubbishData).Updates(model.MapMainToModel(updated))

	data = model.MapModelToMain(rubbishData)
	return data, nil
}
