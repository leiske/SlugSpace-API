package models

type Lot struct {

	Id int64 `json:"id"`

	Name string `json:"name"`

	FreeSpaces int64 `json:"freeSpaces"`

	TotalSpaces int64 `json:"totalSpaces"`

	LastUpdated string `json:"lastUpdated"`
}
