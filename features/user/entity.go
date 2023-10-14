package user

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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type UserDataInterface interface {
	Create(data Main) error
	CheckLogin(email, password string) (Main, string, error)
}

type UseCaseInterface interface {
	Create(data Main) error
	CheckLogin(email, password string) (Main, string, error)
}
