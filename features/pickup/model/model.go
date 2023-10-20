package model

import (
	"recycle/features/location/model"
	"time"

	"gorm.io/gorm"
)

type Pickup struct {
	Id         string `gorm:"type:varchar(100);primaryKey;not null" json:"id"`
	Address    string `gorm:"type:longtext;not null" json:"address"`
	Longitude  string `gorm:"type:varchar(255);not null" json:"longitude"`
	Latitude   string `gorm:"type:varchar(255);not null" json:"latitude"`
	PickupDate string `gorm:"type:date;not null" json:"pickup_date"`
	Status     string `gorm:"type:varchar(20);not null" json:"status"`
	UserId     string `gorm:"type:varchar(100)" json:"user_id"`
	LocationId string `json:"location_id"`
	Location   model.Location
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
