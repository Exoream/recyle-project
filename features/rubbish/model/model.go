package model

import (
	"time"

	"gorm.io/gorm"
)

type Rubbish struct {
	Id          string         `gorm:"type:varchar(100);primaryKey;not null" json:"id"`
	Name        string         `gorm:"varchar(255);not null;unique" json:"name"`
	TypeRubbish string         `gorm:"varchar(100);not null" json:"type_rubbish"`
	PointPerKg  int            `gorm:"not null" json:"point_per_kg"`
	Description string         `gorm:"type:longtext" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}