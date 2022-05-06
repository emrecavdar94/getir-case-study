package record

import "time"

type RecordRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type Record struct {
	Key       string    `json:"key" bson:"key"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Counts    int       `json:"counts" bson:"totalCount"`
}

type RecordDTO struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int    `json:"totalCount"`
}
