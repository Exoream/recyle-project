package rubbish

import "time"

type Main struct {
	Id          string
	Name        string `validate:"required"`
	TypeRubbish string `validate:"required"`
	PointPerKg  int    `validate:"required"`
	Description string 
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}


type RubbishDataInterface interface {
	Create(data Main) error
	GetById(id string) (Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllRubbish() ([]Main, error)
}

type UseCaseInterface interface {
	Create(data Main) error
	GetById(id string) (Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllRubbish() ([]Main, error)
}