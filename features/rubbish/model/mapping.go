package model

import "recycle/features/rubbish"

// Mapping dari Main ke Model
func MapMainToModel(mainData rubbish.Main) Rubbish {
	return Rubbish{
		Name:        mainData.Name,
		TypeRubbish: mainData.TypeRubbish,
		PointPerKg:  mainData.PointPerKg,
		Description: mainData.Description,
	}
}

// Mapping dari Model ke Main
func MapModelToMain(mainData Rubbish) rubbish.Main {
	return rubbish.Main{
		Id:          mainData.Id,
		Name:        mainData.Name,
		TypeRubbish: mainData.TypeRubbish,
		PointPerKg:  mainData.PointPerKg,
		Description: mainData.Description,
		CreatedAt:   mainData.CreatedAt,
		UpdatedAt:   mainData.UpdatedAt,
	}
}

func ModelToMainMapping(dataModel []Rubbish) []rubbish.Main {
	var mainList []rubbish.Main
	for _, value := range dataModel {
		mainList = append(mainList, MapModelToMain(value))
	}
	return mainList
}
