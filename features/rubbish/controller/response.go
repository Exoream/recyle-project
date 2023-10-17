package controller

import (
	"recycle/features/rubbish"
)

type RubbishRespon struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	TypeRubbish string    `json:"type_rubbish"`
	PointPerKg  int       `json:"point_per_kg"`
	Description string    `json:"description"`
}

func MainResponse(dataMain rubbish.Main) RubbishRespon {
	return RubbishRespon{
		Id:          dataMain.Id,
		Name:        dataMain.Name,
		TypeRubbish: dataMain.TypeRubbish,
		PointPerKg:  dataMain.PointPerKg,
		Description: dataMain.Description,
	}
}
