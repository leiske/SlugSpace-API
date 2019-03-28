package models

type LotPredictedFreespace struct {
	ID int `json:"id"`

	PredictedFreespace int `json:"predictedFreespace"`

	IsPredicted bool `json:"IsPredicted"`
}
