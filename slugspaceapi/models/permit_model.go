package models

type Permit struct {
	Id int64 `json:"id"`

	PermitName string `json:"permitName"`

	PermitInfo string `json:"permitInfo"`
}
