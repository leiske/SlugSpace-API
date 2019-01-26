package models

type UntrackedLot struct {
	Id int64 `json:"id"`

	LotName string `json:"lotName"`

	LotNumber int64 `json:"lotNumber"`

	Longitude float64 `json:"longitude"`

	Latitude float64 `json:"latitude"`

	Permits string `json:"permits"`

	FreeAfter string `json:"freeAfter"`

	PayStations string `json:"payStations"`

}
