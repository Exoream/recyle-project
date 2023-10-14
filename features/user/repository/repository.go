package repository

import (
	"errors"
	"recycle/app/middlewares"
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
func (u *userRepository) CheckLogin(email string, password string) (user.Main, string, error) {
	var data model.User

	tx := u.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return user.Main{}, "", tx.Error
	}

	if helper.CheckPasswordHash(data.Password, password) {
		id, err := uuid.Parse(data.ID)
		if err != nil {
			return user.Main{}, "", err
		}

		token, errToken := middlewares.CreateToken(id)
		if errToken != nil {
			return user.Main{}, "", errToken
		}

		dataMain := model.MapModelToMain(data)
		return dataMain, token, nil
	}

	return user.Main{}, "", errors.New("Login failed")
}
