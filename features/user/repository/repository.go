package repository

import (
	"errors"
	"recycle/features/user"
	"recycle/features/user/model"
	"recycle/helper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserDataInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.UserDataInterface.
func (u *userRepository) Create(user user.Main) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	hashedPassword, errHash := helper.HashPassword(user.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}

	dataInput := model.MapMainToModel(user)
	dataInput.ID = newUUID.String()
	dataInput.Password = hashedPassword

	tx := u.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// CheckLogin implements user.UserDataInterface.
func (u *userRepository) CheckLogin(email string, password string) (user.Main, error) {
	var data model.User

    tx := u.db.Where("email = ?", email).First(&data)
    if tx.Error != nil {
        return user.Main{}, tx.Error
    }

    dataMain := model.MapModelToMain(data)
    return dataMain, nil
}

// GetById implements user.UserDataInterface.
func (u *userRepository) GetById(id string) (user.Main, error) {
    var userData model.User
    result := u.db.Where("id = ?", id).First(&userData)
    if result.Error != nil {
        return user.Main{}, result.Error
    }

    var userById = model.MapModelToMain(userData)
    return userById, nil
}

// UpdateById implements user.UserDataInterface.
func (u *userRepository) UpdateById(id string, updated user.Main) (data user.Main, err error) {
	uuidID, err := uuid.Parse(id)
    if err != nil {
        return user.Main{}, err
    }

    var usersData model.User
    resultFind := u.db.First(&usersData, uuidID)
    if resultFind.Error != nil {
        return user.Main{}, resultFind.Error
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
func (u *userRepository) FindAllUsers() ([]user.Main, error) {
	var users []model.User

	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var allUser []user.Main = model.ModelToMainMapping(users)

	return allUser, nil
}