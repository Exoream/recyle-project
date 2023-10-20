package entity

import "time"

type Main struct {
	Id         string
	Address    string `validate:"required"`
	Longitude  string `validate:"required"`
	Latitude   string `validate:"required"`
	PickupDate string `validate:"required"`
	Status     string
	UserId     string
	LocationId string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type PickupDataInterface interface {
	Create(data Main) error
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	GetById(id string) (Main, error)
	FindAllPickup() ([]Main, error)
}

type UseCaseInterface interface {
	Create(data Main) error
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	GetById(id string) (Main, error)
	FindAllPickup() ([]Main, error)
}
