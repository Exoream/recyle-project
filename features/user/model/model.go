package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"type:varchar(100);primaryKey;not null" json:"id"`
	Name        string         `gorm:"varchar(255);not null" json:"name"`
	Email       string         `gorm:"varchar(255);unique;not null" json:"email"`
	Password    string         `gorm:"varchar(255);not null" json:"password"`
	Gender      string         `gorm:"type:enum('Man','Woman')" json:"gender"`
	Age         int            `json:"age"`
	Address     string         `gorm:"type:longtext" json:"address"`
	SaldoPoints int            `json:"saldo_points"`
	Role        string         `json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
