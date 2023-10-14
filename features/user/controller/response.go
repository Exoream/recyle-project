package controller

type UserResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	SaldoPoints int    `json:"saldo_points"`
}

type UserLoginResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


