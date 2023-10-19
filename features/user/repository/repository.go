package repository

import (
	"errors"
	"recycle/features/user/entity"
	"recycle/features/user/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserDataInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.UserDataInterface.
func (u *userRepository) Create(user entity.Main) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	dataInput := model.MapMainToModel(user)
	dataInput.ID = newUUID.String()

	tx := u.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// CheckLogin implements user.UserDataInterface.
func (u *userRepository) CheckLogin(email string, password string) (entity.Main, error) {
	var data model.User

    tx := u.db.Where("email = ?", email).First(&data)
    if tx.Error != nil {
        return entity.Main{}, tx.Error
    }

    dataMain := model.MapModelToMain(data)
    return dataMain, nil
}

// GetById implements user.UserDataInterface.
func (u *userRepository) GetById(id string) (entity.Main, error) {
    var userData model.User
    result := u.db.Where("id = ?", id).First(&userData)
    if result.Error != nil {
        return entity.Main{}, result.Error
    }

    var userById = model.MapModelToMain(userData)
    return userById, nil
}

// UpdateById implements user.UserDataInterface.
func (u *userRepository) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	uuidID, err := uuid.Parse(id)
    if err != nil {
        return entity.Main{}, err
    }

    var usersData model.User
    resultFind := u.db.First(&usersData, uuidID)
    if resultFind.Error != nil {
        return entity.Main{}, resultFind.Error
    }

    u.db.Model(&usersData).Updates(model.MapMainToModel(updated))

    data = model.MapModelToMain(usersData)
    return data, nil
}

// Delete implements user.UserDataInterface.
func (u *userRepository) DeleteById(id string) error {
	uuidID, err := uuid.Parse(id)
    if err != nil {
        return err
    }

	var user model.User
	result := u.db.Where("id = ?", uuidID).Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete data, record not found")
	}

	return nil
}

// FindAllUsers implements user.UserDataInterface.
func (u *userRepository) FindAllUsers() ([]entity.Main, error) {
	var users []model.User

	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var allUser []entity.Main = model.ModelToMainMapping(users)

	return allUser, nil
}