package models

type Lot struct {
	Id int64 `json:"id"`

	FullName string `json:"fullName"`

	Name string `json:"name"`

	Description string `json:"description"`

	ImageURI string `json:"imageURI"`

	FreeSpaces int64 `json:"freeSpaces"`

	TotalSpaces int64 `json:"totalSpaces"`

	Longitude float64 `json:"longitude"`

	Latitude float64 `json:"latitude"`

	LastUpdated string `json:"lastUpdated"`
}
