package entity

import "time"

type Main struct {
	Id          string
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Gender      string `validate:"oneof=Man Woman"`
	Age         int
	Address     string
	SaldoPoints int
	Role		string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type UserDataInterface interface {
	Create(data Main) error
	CheckLogin(email, password string) (Main, error)
	GetById(id string) (Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllUsers() ([]Main, error)
}

type UseCaseInterface interface {
	Create(data Main) error
	CheckLogin(email, password string) (Main, string, error)
	GetById(id string) (Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllUsers() ([]Main, error)
}
