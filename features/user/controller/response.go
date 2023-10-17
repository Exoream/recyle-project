package controller

import (
	"recycle/features/user"
	"time"
)


type UserLoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
	SaldoPoints int       `json:"saldo_points"`
	CreatedAt   time.Time `json:"created_at"`
}

func MainResponse(dataMain user.Main) UserResponse {
	return UserResponse{
		Id:          dataMain.Id,
		Name:        dataMain.Name,
		Email:       dataMain.Email,
		Gender:      dataMain.Gender,
		Age:         dataMain.Age,
		Address:     dataMain.Address,
		SaldoPoints: dataMain.SaldoPoints,
		CreatedAt:   dataMain.CreatedAt,
	}
}

func LoginResponse(id, email, token string) UserLoginResponse {
    return UserLoginResponse{
        Id:    id,
        Email: email,
        Token: token,
    }
}


