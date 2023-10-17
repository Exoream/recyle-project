package controller

import "recycle/features/rubbish"

type RubbishRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	TypeRubbish string `json:"type_rubbish"`
	PointPerKg  int    `json:"point_per_kg"`
	Description string `json:"description"`
}

func RequestMain(dataRequest RubbishRequest) rubbish.Main {
	return rubbish.Main{
		Id:          dataRequest.Id,
		Name:        dataRequest.Name,
		TypeRubbish: dataRequest.TypeRubbish,
		PointPerKg:  dataRequest.PointPerKg,
		Description: dataRequest.Description,
	}
}
