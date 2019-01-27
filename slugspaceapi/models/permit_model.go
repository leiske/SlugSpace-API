package models

type Permit struct {
	Id int64 `json:"id"`

	Name string `json:"permitName"`

	Info string `json:"permitInfo"`
}
