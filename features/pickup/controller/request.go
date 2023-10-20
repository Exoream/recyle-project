package controller

import "recycle/features/pickup/entity"

type PickupRequest struct {
	Id         string `json:"id"`
	Address    string `json:"address"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	PickupDate string `json:"pickup_date"`
	UserId     string `json:"user_id"`
	LocationId string `json:"location_id"`
}

func RequestMain(dataRequest PickupRequest) entity.Main {
	return entity.Main{
		Id:         dataRequest.Id,
		Address:    dataRequest.Address,
		Longitude:  dataRequest.Longitude,
		Latitude:   dataRequest.Latitude,
		PickupDate: dataRequest.PickupDate,
		LocationId: dataRequest.LocationId,
	}
}
