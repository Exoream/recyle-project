package model

import "recycle/features/user"

// Mapping dari Main ke Model
func MappingModel(mainData user.Main) User {
	return User{
		Name:        mainData.Name,
		Email:       mainData.Email,
		Password:    mainData.Password,
		Gender:      mainData.Gender,
		Age:         mainData.Age,
		Address:     mainData.Address,
		SaldoPoints: mainData.SaldoPoints,
	}
}

// // Mapping dari Model ke Main
// func MappingMain(mainData User) user.Main {
// 	return user.Main{
// 		Id:          mainData.ID,
// 		Name:        mainData.Name,
// 		Email:       mainData.Email,
// 		Password:    mainData.Password,
// 		Gender:      mainData.Gender,
// 		Age:         mainData.Age,
// 		Address:     mainData.Address,
// 		SaldoPoints: mainData.SaldoPoints,
// 		CreatedAt:   mainData.CreatedAt,
// 		UpdatedAt:   mainData.UpdatedAt,
// 	}
// }

// func ModelToCoreMapping(dataModel []User) []user.Main {
// 	var coreList []user.Main
// 	for _, v := range dataModel {
// 		coreList = append(coreList, MappingMain(v))
// 	}
// 	return coreList
// }
