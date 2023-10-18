package controller

import "recycle/features/rubbish/entity"

type RubbishRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	TypeRubbish string `json:"type_rubbish"`
	PointPerKg  int    `json:"point_per_kg"`
	Description string `json:"description"`
}

func RequestMain(dataRequest RubbishRequest) entity.Main {
	return entity.Main{
		Id:          dataRequest.Id,
		Name:        dataRequest.Name,
		TypeRubbish: dataRequest.TypeRubbish,
		PointPerKg:  dataRequest.PointPerKg,
		Description: dataRequest.Description,
	}
}
