package model

import (
	"recycle/features/pickup/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          string         `gorm:"type:varchar(100);primaryKey;not null" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Email       string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password    string         `gorm:"type:varchar(255);not null" json:"password"`
	Gender      string         `gorm:"type:enum('Man','Woman')" json:"gender"`
	Age         int            `json:"age"`
	Address     string         `gorm:"type:longtext" json:"address"`
	SaldoPoints float64        `json:"saldo_points"`
	Role        string         `json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Pickups     []model.Pickup
}
