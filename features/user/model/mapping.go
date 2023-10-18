package model

import "recycle/features/user/entity"

// Mapping dari Main ke Model
func MapMainToModel(mainData entity.Main) User {
	return User{
		Name:        mainData.Name,
		Email:       mainData.Email,
		Password:    mainData.Password,
		Gender:      mainData.Gender,
		Age:         mainData.Age,
		Address:     mainData.Address,
		SaldoPoints: mainData.SaldoPoints,
		Role:        mainData.Role,
	}
}

// Mapping dari Model ke Main
func MapModelToMain(mainData User) entity.Main {
	return entity.Main{
		Id:          mainData.ID,
		Name:        mainData.Name,
		Email:       mainData.Email,
		Password:    mainData.Password,
		Gender:      mainData.Gender,
		Age:         mainData.Age,
		Address:     mainData.Address,
		SaldoPoints: mainData.SaldoPoints,
		Role:        mainData.Role,
		CreatedAt:   mainData.CreatedAt,
		UpdatedAt:   mainData.UpdatedAt,
	}
}

func ModelToMainMapping(dataModel []User) []entity.Main {
	var mainList []entity.Main
	for _, value := range dataModel {
		mainList = append(mainList, MapModelToMain(value))
	}
	return mainList
}
