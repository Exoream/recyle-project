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

	dataInput := model.MappingModel(user)
	dataInput.ID = newUUID.String()
	dataInput.Password = hashedPassword

	tx := u.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
